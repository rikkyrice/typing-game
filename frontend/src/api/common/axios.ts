import axios from 'axios';

export default axios.create({
  timeout: 30000,
  withCredentials: !(
    process.env.VUE_APP_MODE === 'local' &&
    process.env.VUE_APP_NOAUTH === 'true'
  ),
  headers: {
    'Content-Type': 'application/json',
  },
});
