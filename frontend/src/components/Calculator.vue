<template>
  <v-container>
    <h3 class="ma-4 blue-grey--text">Parameters :</h3>

    <v-row class="ma-2">
      <v-col cols="12">
        <v-slider
          label="Total Value"
          v-model="totalValue"
          step="0.1"
          thumb-label="always"
          class="align-center"
          :max="10"
          :min="1"
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
      <v-col cols="8">
        <v-slider
          label="Equity"
          v-model="equity"
          step="0.1"
          thumb-label="always"
          class="align-center"
          :max="5"
          :min="0"
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
      <v-col cols="6">
        <v-slider
          label="Interest Rate"
          v-model="intersetRate"
          step="0.1"
          thumb-label="always"
          class="align-center"
          :max="5"
          :min="1"
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
    </v-row>

    <v-row class="ma-2">
      <v-col cols="6">
        <v-slider
          label="Repayment Period"
          v-model="repaymentReriod"
          step="1"
          thumb-label="always"
          class="align-center"
          :max="30"
          :min="1"
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
    </v-row>

    <v-row class="ma-2">
      <v-spacer></v-spacer>
      <v-col cols="2">
        <v-btn color="success" @click="send">Get Result</v-btn>
      </v-col>
      <v-spacer></v-spacer>
    </v-row>

    <v-row class="ma-2">
      <h3 class="ma-2 mt-8 blue-grey--text">{{ 'result : ' + result }}</h3>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue';
import { setParameters, ParameterData } from '../services/data';

export default Vue.extend({
  components: {},
  data(): {
    totalValue: number;
    equity: number;
    result: number;
    loan: string;
    intersetRate: number;
    repaymentReriod: number;
  } {
    return {
      totalValue: 3,
      equity: 1,
      result: 0,
      loan: '2',
      intersetRate: 2.4,
      repaymentReriod: 10,
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
        houseValue: this.totalValue,
        equity: this.equity,
        interestRate: this.intersetRate,
        paymentPeriod: this.repaymentReriod,
        type: 'loan',
      };
      setParameters(data)
        .then((response) => this.setResult(response))
        .catch((error) => {
          console.log(error);
        });
    },

    //server got the parameter
    setResult(data: any) {
      this.result = 0;
    },
  },
});
</script>
