import { z } from 'zod';
export const propertySchema = z.object({
	name: z
		.string()
		.min(2, 'Property name cannot be less than 2 characters')
		.max(50, 'Property name cannot exceed 50 characters'),
	location: z
		.string({ required_error: "The property's location is required" })
		.min(2, 'Property cannot be less than 2 characters'),
	mainRoad: z.string().min(2, 'Main road cannot be less than 2 characters')
});
export type PropertySchema = typeof propertySchema;
