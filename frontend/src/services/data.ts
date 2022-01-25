import Axios from 'axios';

export interface ParameterData {
  houseValue: number;
  equity: number;
  interestRate: number;
  paymentPeriod: number;
  type?: string;
}

const http = Axios.create({
  baseURL: process.env.VUE_APP_API_BASE_PATH,
});

http.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (
      error.response &&
      error.response.status >= 400 &&
      error.response.status < 500
    ) {
      console.log('Logging the error', error);
    }
    throw error;
  }
);

export async function setParameters(data: ParameterData): Promise<void> {
  return http
    .post<void>(`/parameters`, data, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then((response) => {
      return response.data;
    });
}
