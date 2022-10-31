from config.config import settings

###
import glob
import os
from torch.utils.data import Dataset
from pathlib import Path

###
import logging
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

formatter = logging.Formatter('%(message)s')
file_handler = logging.FileHandler('test.log',encoding=settings.encoding)
file_handler.setFormatter(formatter)

### Import pre-trained ML model and related libraries ###
from joblib import load
#from sklearn.linear_model import LogisticRegression
import pandas as pd
###
###Text preprocessing###
import re # Removing punctutions and symbols
import nltk
from nltk.corpus import stopwords
nltk.download('stopwords') #Stop words removal
nltk.download('punkt')
from nltk.stem.snowball import SnowballStemmer #Stemming
import pymystem3 # Lemmatization
import csv
###

### Import Libraries for word embedding 
import torch
from transformers import AutoTokenizer, AutoModel
###

# Define path to selected Dataset via relative approach

absolute_path = os.path.dirname(__file__)
relative_path = settings.rel_path
full_path = os.path.join(absolute_path, relative_path)

class IntentsDataset(Dataset):

    def __init__(self, data = full_path, action=settings.p_action):
        self.data = data
        self.action = action

    def __len__(self):
        dir_pr = f'{self.data}/**/{self.action}*.txt'
        dir_str = glob.glob(dir_pr, recursive=True)
        dicts = []
        for value in dir_str:
            open_txt = open(value, 'r', encoding=settings.encoding)
            txt = open_txt.read()
            dicts.append(txt)
        return len(dicts)

    def __getitem__(self, id_x):
        dir_pr = f'{self.data}/**/{self.action}*.txt'
        dir_str = glob.glob(dir_pr, recursive=True)
        dicts = {'input':[], 'output':[]}
        dict_to_return = {'input':'', 'output':''}
        for value in dir_str:
            intent_name = os.path.basename(str(Path(value).parents[1]))
            open_txt = open(value, 'r', encoding=settings.encoding)
            txt = open_txt.read()
            dicts['input'].append(txt)
            dicts['output'].append(intent_name)
        dict_to_return['input']=dicts['input'][id_x]
        dict_to_return['output']=dicts['output'][id_x]
        return dict_to_return

class TextPreprocessor:
        
    def lower(self, string): # lowercase
        return string.lower()
    
    def rem_punc_symb(self, string):#Removing punctutions and symbols
        return re.sub(r'[^\w\s]','',string)
    
    def tokenize(self, string):#Tokenization
        return nltk.word_tokenize(string)
    
    def stopword_rem(self, string):#Stop words removal
        stopwordss = stopwords.words(settings.prs_language)
        strings = self.tokenize(string)
        f_w = []
        for word in strings:
            if (word not in stopwordss):
                f_w.append(word)
        return ' '.join(f_w)
    
    def stemming(self, string): #Stemming
        stemmer = SnowballStemmer(settings.prs_language)
        input_str = self.tokenize(string)
        return ' '.join([stemmer.stem(word) for word in input_str])
    
    def lemma(self, string): #Lemmatization
        lemmatizer = pymystem3.Mystem()
        input_str = self.tokenize(string)
        output_str = []
        for word in input_str:
            output_str.append(lemmatizer.lemmatize(word))
        return ' '.join([i[0] for i in output_str]).replace('\n', '')
    
    
    def pipuline_stem(self, string):#Preprocessing till stem
        low = self.lower(string)
        rem_p = self.rem_punc_symb(low)
        stopw = self.stopword_rem(rem_p)
        stem = self.stemming(stopw)
        return stem
    
    def __call__(self, string):#Preprocessing till lemmatization 
        low = self.lower(string)
        rem_p = self.rem_punc_symb(low)
        stopw = self.stopword_rem(rem_p)
        lemmu = self.lemma(stopw)
        return lemmu

    def combo_d(self): # Input after text processing and output as is.
        test = IntentsDataset(full_path)
        txtprs = TextPreprocessor()
        in_to_csv = []
        out_to_csv = []
        for item in range(0, len(test)):
            in_to_csv.append(txtprs((test[item])['input'])) 
            out_to_csv.append(txtprs((test[item])['output'])) # append post-processed text to list
        comb_d = list(zip(in_to_csv,out_to_csv))
        logger.info(f'Data after preprocessing {comb_d[0:3]}...{comb_d[-3:]}')
        return comb_d 

class Embedding:
    def __init__(self):
        self.tokenizer = AutoTokenizer.from_pretrained(settings.rubert)
        self.model = AutoModel.from_pretrained(settings.rubert)

    def embed_bert_cls(self, text):
        t = self.tokenizer(text, padding=True, truncation=True, return_tensors='pt')
        with torch.no_grad():
            model_output = self.model(**{k: v.to(self.model.device) for k, v in t.items()})
        embeddings = model_output.last_hidden_state[:, 0, :]
        embeddings = torch.nn.functional.normalize(embeddings)
        return embeddings[0].cpu().numpy()

# use pre-trained ML model
class PredictSentence:
    def __init__(self):
        self.model = load('final_logr_model.joblib')
    def predict_model(self, sentence):
        text_pr = TextPreprocessor()
        text_af_pr = text_pr(sentence)
        embed_txt = Embedding()
        embed_txt_pr = embed_txt.embed_bert_cls(text_af_pr)
        datum = pd.DataFrame(embed_txt_pr).T
        pred = self.model.predict(datum)[0]
        logger.info(f'Input: {sentence} Output: {pred}') 
        return pred


