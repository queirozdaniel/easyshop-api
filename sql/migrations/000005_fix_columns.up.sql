ALTER TABLE public.users_resources ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE public.users_resources ALTER COLUMN role_resource_id SET NOT NULL;

ALTER TABLE public.roles_resources ALTER COLUMN resource_id SET NOT NULL;
ALTER TABLE public.roles_resources ALTER COLUMN role_id SET NOT NULL;

ALTER TABLE public.resources ALTER COLUMN "name" SET NOT NULL;

ALTER TABLE public.roles ALTER COLUMN "name" SET NOT NULL;

ALTER TABLE public.persons ALTER COLUMN user_id SET NOT NULL;

ALTER TABLE public.users ALTER COLUMN "password" SET NOT NULL;
ALTER TABLE public.users ALTER COLUMN email SET NOT NULL;