import Axios from 'axios';

export interface ParameterData {
  client: string;
  houseValue: number;
  equity: number;
  interestRate: number;
  paymentYear: number;
  oneTimeFee: number;
  periodicFee: number;
  type?: string;
}

export interface ResultData {
  client: string;
  totalInterest: number;
  periodicPayment: number;
  totalPayment: number;
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

export async function setParameters(data: ParameterData): Promise<ResultData> {
  return http
    .post<ResultData>(`/parameters`, data, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then((response) => {
      console.log(response.data);
      return response.data;
    });
}
