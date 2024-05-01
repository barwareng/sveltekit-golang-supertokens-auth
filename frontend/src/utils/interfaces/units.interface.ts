export interface IUnit {
	id: string;
	propertyId: string;
	label: string;
	rent: number;
	rooms: number;
	baths: number;
	imageUrls: string[];
	areaSqm: number;
	areaSqf: number;
	internalFeatures: string[];
	externalFeatures: string[];
	status: 'LISTED' | 'OCCUPIED' | 'VACANT' | 'IN ARREARS';
}
