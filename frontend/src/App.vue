<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-img
          alt="Money Logo"
          class="shrink mr-2"
          contain
          src="./assets/images/dollar.png"
          transition="scale-transition"
          width="50"
        />
        <h3 class="ml-2">Investment Calculator</h3>
      </div>
    </v-app-bar>
    <v-main>
      <calculator />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import Calculator from './components/Calculator.vue';

export default Vue.extend({
  name: 'App',

  components: {
    Calculator,
  },

  data: () => ({
    //
  }),
});
</script>
