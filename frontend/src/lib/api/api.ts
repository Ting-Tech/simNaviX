const BASE_URL = 'http://localhost:8080';

export async function getRobot() {
	const res = await fetch(`${BASE_URL}/robot`);
	if (!res.ok) {
		throw new Error('Failed to fetch robot data');
	}
	return await res.json();
}
