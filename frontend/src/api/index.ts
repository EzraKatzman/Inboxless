import axios from 'axios';

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL; 

export async function createInbox() {
  const response = await axios.post(`${BASE_URL}/api/inbox`);
  return response.data;
}
