-- Table: public.produtos

-- DROP TABLE public.produtos;

CREATE TABLE IF NOT EXISTS public.produtos
(
    id integer NOT NULL DEFAULT nextval('produtos_id_seq'::regclass),
    nome character varying COLLATE pg_catalog."default",
    descricao character varying COLLATE pg_catalog."default",
    preco numeric,
    quantidade integer,
    CONSTRAINT produtos_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.produtos
    OWNER to postgres;