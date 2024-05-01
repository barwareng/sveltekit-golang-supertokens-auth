import { z } from 'zod';
const MAX_FILE_SIZE = 10000000;
const ACCEPTED_IMAGE_TYPES = ['image/jpeg', 'image/jpg', 'image/png', 'image/webp'];
export const teamSchema = z.object({
	name: z
		.string()
		.min(2, 'Name cannot be less than 2 characters')
		.max(50, 'Name cannot exceed 50 characters'),
	description: z
		.string()
		.min(10, 'Description cannot be less than 10 characters')
		.max(300, 'Description cannot exceed 300 characters')
		.optional(),
	coverImage: z
		.any()
		.refine((file) => file?.size <= MAX_FILE_SIZE, `Max image size is 5MB.`)
		.refine(
			(file) => ACCEPTED_IMAGE_TYPES.includes(file?.type),
			'Only .jpg, .jpeg, .png and .webp formats are supported.'
		)
		.optional()
});
export type TeamSchema = typeof teamSchema;
