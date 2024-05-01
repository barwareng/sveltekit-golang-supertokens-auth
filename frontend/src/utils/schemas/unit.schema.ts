import { z } from 'zod';
export const unitSchema = z.object({
	propertyId: z.string({ required_error: 'Please select a property' }),
	label: z.string({ required_error: 'Please select a property' }),
	rent: z.coerce.number({ required_error: "Please enter this unit's monthly rent" }),
	rooms: z.coerce.number({ required_error: 'Please enter the number of bedrooms in this unit' }),
	baths: z.coerce.number({ required_error: 'Please enter the number of bathrooms in this unit' }),
	imageUrls: z.string().array().max(10, 'You can choose up to ten images.').optional(),
	areaSqm: z.coerce.number().optional(),
	areaSqf: z.coerce.number().optional(),
	internalFeatures: z.string().array().optional(),
	externalFeatures: z.string().array().optional()
});
export type UnitSchema = typeof unitSchema;
