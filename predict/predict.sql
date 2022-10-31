CREATE TABLE public.predicts (
	id serial NOT NULL,
	sentence varchar(64) NOT NULL,
	intent varchar(256) NULL,
	user_id int4 NULL,
	CONSTRAINT predicts_pkey PRIMARY KEY (id)
);