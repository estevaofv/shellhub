import http from '@/helpers/http';

const postUser = async (user) => http().post('/userTeste', { user });

export { postUser as default };
