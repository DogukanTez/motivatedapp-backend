CREATE TABLE public.checkins (
	id uuid NOT NULL,
	goal_id uuid NOT NULL,
	checkin_date date NOT NULL,
	status text NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT checkins_goal_id_checkin_date_key UNIQUE (goal_id, checkin_date),
	CONSTRAINT checkins_pkey PRIMARY KEY (id),
	CONSTRAINT checkins_status_check CHECK ((status = ANY (ARRAY['success'::text, 'fail'::text])))
);


-- public.checkins foreign keys

ALTER TABLE public.checkins ADD CONSTRAINT checkins_goal_id_fkey FOREIGN KEY (goal_id) REFERENCES public.goals(id) ON DELETE CASCADE;

CREATE TABLE public.goals (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	title text NOT NULL,
	"type" text NOT NULL,
	start_date date NOT NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT goals_pkey PRIMARY KEY (id),
	CONSTRAINT goals_type_check CHECK ((type = ANY (ARRAY['do'::text, 'avoid'::text])))
);


-- public.goals foreign keys

ALTER TABLE public.goals ADD CONSTRAINT goals_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


CREATE TABLE public.subscriptions (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	platform text NOT NULL,
	product_id text NOT NULL,
	is_lifetime bool DEFAULT false NOT NULL,
	started_at timestamptz NOT NULL,
	expires_at timestamptz NULL,
	status text NOT NULL,
	original_transaction_id text NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT subscriptions_pkey PRIMARY KEY (id),
	CONSTRAINT subscriptions_platform_check CHECK ((platform = ANY (ARRAY['ios'::text, 'android'::text]))),
	CONSTRAINT subscriptions_status_check CHECK ((status = ANY (ARRAY['active'::text, 'expired'::text, 'canceled'::text])))
);
CREATE UNIQUE INDEX unique_active_subscription ON public.subscriptions USING btree (user_id) WHERE (status = 'active'::text);


-- public.subscriptions foreign keys

ALTER TABLE public.subscriptions ADD CONSTRAINT subscriptions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;



CREATE TABLE public.user_details (
	user_id uuid NOT NULL,
	device_id text NULL,
	platform text NULL,
	locale text NULL,
	timezone text NULL,
	notification_time time NULL,
	last_seen_at timestamptz NULL,
	CONSTRAINT user_details_pkey PRIMARY KEY (user_id),
	CONSTRAINT user_details_platform_check CHECK ((platform = ANY (ARRAY['ios'::text, 'android'::text])))
);


-- public.user_details foreign keys

ALTER TABLE public.user_details ADD CONSTRAINT user_details_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


CREATE TABLE public.users (
	id uuid NOT NULL,
	email text NULL,
	password_hash text NULL,
	auth_provider text DEFAULT 'local'::text NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	updated_at timestamptz NULL,
	CONSTRAINT users_auth_provider_check CHECK ((auth_provider = ANY (ARRAY['local'::text, 'apple'::text, 'google'::text, 'anonymous'::text]))),
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);