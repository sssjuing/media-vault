import axios from 'axios';

export async function requestLogin(values: { username: string; password: string }) {
  await axios.post('/api/user/login', values);
}
