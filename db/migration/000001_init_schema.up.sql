CREATE TABLE
  public.countries (
   	id serial NOT NULL,
    name character varying(255) NOT NULL,
    shortname character varying(255) NOT NULL,
    continent character varying(255) NOT NULL,
    is_operational boolean NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
  );

ALTER TABLE
  public.countries
ADD
  CONSTRAINT countries_table_pkey PRIMARY KEY (id)