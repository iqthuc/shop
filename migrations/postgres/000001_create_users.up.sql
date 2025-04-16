-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    username character varying(32) COLLATE pg_catalog."default",
    email character varying(128) COLLATE pg_catalog."default" NOT NULL,
    password_hash character varying(64) COLLATE pg_catalog."default",
    role character varying(16) COLLATE pg_catalog."default" DEFAULT 'user'::text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_username_key UNIQUE (username)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to root;

GRANT ALL ON TABLE public.users TO iqthuc WITH GRANT OPTION;

GRANT ALL ON TABLE public.users TO root;
-- Index: users_email_idx

-- DROP INDEX IF EXISTS public.users_email_idx;

CREATE INDEX IF NOT EXISTS users_email_idx
    ON public.users USING btree
    (email COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: users_username_idx

-- DROP INDEX IF EXISTS public.users_username_idx;

CREATE INDEX IF NOT EXISTS users_username_idx
    ON public.users USING btree
    (username COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
