CREATE TABLE IF NOT EXISTS eps (
                     nit int4 NOT NULL,
                     nombre varchar NOT NULL,
                     activa bool NULL,
                     CONSTRAINT eps_pk PRIMARY KEY (nit)
);

CREATE TABLE IF NOT EXISTS tipo_documento (
                                id int4 NOT NULL,
                                detalle varchar NOT NULL,
                                CONSTRAINT tipo_documento_pk PRIMARY KEY (id)
);



CREATE TABLE IF NOT EXISTS tipo_sexo (
                                        id int4 NOT NULL,
                                        detalle varchar NOT NULL,
                                        CONSTRAINT tipo_sexo_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
                       id serial4 NOT NULL,
                       firstname varchar(255) NOT NULL,
                       lastname varchar(255) NOT NULL,
                       email varchar(255) NOT NULL,
                       "password" varchar(255) NOT NULL,
                       created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       CONSTRAINT users_pkey PRIMARY KEY (id)
);
