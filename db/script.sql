ALTER TABLE public.merchant_group DROP COLUMN portal_status;
ALTER TABLE public.merchant_group ADD portal_status int NULL DEFAULT 0;
COMMENT ON COLUMN public.merchant_group.portal_status IS '0 = Inactive, 1 = Active';