<template>
  <v-container>
    <h3 class="ma-4 blue-grey--text">Client :</h3>
    <v-row class="ma-1">
      <v-col cols="4">
        <v-text-field v-model="client" label="Name"></v-text-field>
      </v-col>
    </v-row>
    <h3 class="ma-4 blue-grey--text">Parameters :</h3>

    <v-row class="ma-2">
      <v-col cols="12">
        <v-slider
          label="Total Value"
          v-model="totalValue"
          step="0.1"
          thumb-label="always"
          class="align-center"
          max="10"
          min="1"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="totalValue"
              label="Amount"
              class="mt-0 pt-0"
              hide-details
              single-line
              step="0.1"
              type="number"
              suffix="million ( 1,000,000 ) kr"
              style="width: 220px"
            ></v-text-field>
          </template>
        </v-slider>
      </v-col>
    </v-row>

    <v-row class="ma-2">
      <v-col cols="12">
        <v-slider
          label="Equity"
          v-model="equity"
          step="0.1"
          thumb-label="always"
          class="align-center"
          max="5"
          min="0"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="equity"
              class="mt-0 pt-0"
              hide-details
              single-line
              step="0.1"
              type="number"
              suffix="million ( 1,000,000 ) kr"
              style="width: 220px"
            ></v-text-field>
          </template>
        </v-slider>
      </v-col>
    </v-row>
    <v-row class="ma-2 mb-8">
      <v-col cols="6">
        <span class="subheading font-weight-light mr-2">Loan amount: </span>
        <span
          class="text-h3 orange--text font-weight-light mr-1"
          v-text="loan"
        ></span>
        <span class="subheading font-weight-light mr-1">million nok</span>
      </v-col>
    </v-row>
    <v-row class="ma-2">
      <v-col cols="8">
        <v-slider
          label="Interest Rate"
          v-model="intersetRate"
          step="0.1"
          thumb-label="always"
          class="align-center"
          max="5"
          min="1"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="intersetRate"
              class="mt-0 pt-0"
              hide-details
              single-line
              step="0.1"
              type="number"
              suffix="%"
              style="width: 100px"
            ></v-text-field>
          </template>
        </v-slider>
      </v-col>
      <v-col cols="4">
        <v-text-field
          class="ml-6 pt-0"
          label="One-Time Fee"
          v-model="oneTimeFee"
          step="100"
          type="number"
          max="10000"
          min="0"
          outlined
        ></v-text-field>
      </v-col>
    </v-row>

    <v-row class="ma-2">
      <v-col cols="8">
        <v-slider
          label="Repayment Year"
          v-model="repaymentReriod"
          step="1"
          thumb-label="always"
          class="align-center"
          max="30"
          min="1"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="repaymentReriod"
              label="Amount"
              class="mt-0 pt-0"
              hide-details
              single-line
              step="1"
              type="number"
              suffix="years"
              style="width: 100px"
            ></v-text-field>
          </template>
        </v-slider>
      </v-col>
      <v-col cols="4">
        <v-text-field
          class="ml-6 pt-0"
          label="Periodic Fee"
          v-model="periodicFee"
          step="100"
          type="number"
          max="1000"
          min="0"
          outlined
        ></v-text-field>
      </v-col>
    </v-row>

    <v-row class="ma-2">
      <v-spacer></v-spacer>
      <v-col cols="2">
        <v-btn color="success" @click="send">Get Result</v-btn>
      </v-col>
      <v-spacer></v-spacer>
      <v-col cols="2">
        <v-btn color="indigo" dark @click="get">Get History</v-btn>
      </v-col>
      <v-spacer></v-spacer>
    </v-row>

    <h3 class="ma-4 blue-grey--text">Result :</h3>
    <v-row class="ma-2 mb-8">
      <v-col cols="6">
        <span class="subheading font-weight-light mr-2">Total Interest: </span>
        <span
          class="text-h4 blue--text font-weight-light mr-1"
          v-text="totalInterest"
        ></span>
        <span class="subheading font-weight-light mr-1">nok</span>
      </v-col>
      <v-col cols="6">
        <span class="subheading font-weight-light mr-2">Monthly Payment: </span>
        <span
          class="text-h4 orange--text font-weight-light mr-1"
          v-text="monthlyPayment"
        ></span>
        <span class="subheading font-weight-light mr-1">nok</span>
      </v-col>
    </v-row>
    <v-row class="ma-2 mb-8">
      <v-col cols="6">
        <span class="subheading font-weight-light mr-2">Total Payment: </span>
        <span
          class="text-h4 red--text font-weight-light mr-1"
          v-text="totalPayment"
        ></span>
        <span class="subheading font-weight-light mr-1">nok</span>
      </v-col>
    </v-row>

    <v-div v-if="clientHistoryData">
      <h3 class="ma-4 blue-grey--text">History :</h3>
      <v-row
        class="ma-2 ml-6 mb-6"
        v-for="(report, index) in clientHistoryData.reports"
        :key="index"
      >
        <span class="subheading font-weight-light mr-2">User: </span>
        <span
          class="subheading font-weight-light mr-4 blue--text"
          v-text="report.client"
        ></span>
        <span class="subheading font-weight-light mr-2">Total Interest: </span>
        <span
          class="subheading font-weight-light mr-4 blue--text"
          v-text="report.totalInterest"
        ></span>
        <span class="subheading font-weight-light mr-2">Monthly Payment: </span>
        <span
          class="subheading font-weight-light mr-4 blue--text"
          v-text="report.periodicPayment"
        ></span>
        <span class="subheading font-weight-light mr-2">Total Payment: </span>
        <span
          class="subheading font-weight-light mr-4 blue--text"
          v-text="report.totalPayment"
        ></span>
      </v-row>
    </v-div>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue';
import {
  setParameters,
  ParameterData,
  ResultData,
  HistoryData,
  getHistory,
} from '../services/data';

export default Vue.extend({
  components: {},
  data(): {
    client: string;
    totalValue: number;
    equity: number;
    result: number;
    loan: string;
    intersetRate: number;
    repaymentReriod: number;
    oneTimeFee: number;
    periodicFee: number;
    totalInterest: number;
    monthlyPayment: number;
    totalPayment: number;
    clientHistoryData: HistoryData | null;
  } {
    return {
      client: 'Ellen',
      totalValue: 3,
      equity: 1,
      result: 0,
      loan: '2',
      intersetRate: 2.4,
      repaymentReriod: 10,
      oneTimeFee: 3000,
      periodicFee: 60,
      totalInterest: 0,
      monthlyPayment: 0,
      totalPayment: 0,
      clientHistoryData: null,
    };
  },
  // computed: {
  //   loanAmount(): string {
  //     if (this.totalValue >= this.equity) {
  //       return (this.totalValue - this.equity).toFixed(1).toString();
  //     }
  //     return '0';
  //   },
  // },
  watch: {
    totalValue() {
      if (this.totalValue >= this.equity) {
        this.loan = (this.totalValue - this.equity).toFixed(1).toString();
      } else {
        this.loan = '0';
      }
    },
    equity() {
      if (this.totalValue >= this.equity) {
        this.loan = (this.totalValue - this.equity).toFixed(1).toString();
      } else {
        this.loan = '0';
      }
    },
  },

  methods: {
    //send the parameters
    send() {
      const data: ParameterData = {
        client: this.client,
        houseValue: this.totalValue * 1000000,
        equity: this.equity * 1000000,
        interestRate: this.intersetRate,
        paymentYear: this.repaymentReriod,
        oneTimeFee: this.oneTimeFee,
        periodicFee: this.periodicFee,
        type: 'loan',
      };
      setParameters(data)
        .then((response) => this.setResult(response))
        .catch((error) => {
          console.log(error);
        });
    },

    //server got the parameter
    setResult(data: ResultData) {
      this.totalInterest = data.totalInterest;
      this.monthlyPayment = data.periodicPayment;
      this.totalPayment = data.totalPayment;
    },

    get() {
      getHistory(this.client)
        .then((response) => (this.clientHistoryData = response))
        .catch((error) => {
          console.log(error);
        });
    },
  },
});
</script>
