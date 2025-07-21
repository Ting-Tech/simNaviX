import type { Time } from '@internationalized/date';

export interface Robot {
	id: number;
	length: number;
	width: number;
	filename: string;
	created_at: Time;
}
