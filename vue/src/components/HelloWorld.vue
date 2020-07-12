<template>
  <div class="hello">
    <h1>WASM Arithmatic ^-^</h1>

    <input type="text" v-model="value1" />
    <span>{{operator}}</span>
    <input type="text" v-model="value2" />
    =
    <input type="text" v-model="calcResult" disabled />
    <hr />
    <button @click="add();" id="addButton">Add</button>
    <button @click="sub();" id="subButton">Sub</button>
    <button @click="mul();" id="subButton">Mult</button>
    <button @click="div();" id="subButton">Div</button>
  </div>
</template>

<script>
// /*eslint no-undef: "off"*/
/* global waAdd, waSub, waMulti, waDivi */
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      operator: "?",
      value1: "",
      value2: "",
      calcResult: ""
    }
  },
  created() { },
  mounted() { },
  methods: {
    add() {
      let values = [this.value1, this.value2]

      // // eslint-disable-next-line
      this.calcResult = waAdd(...values)
      this.operator = "+"
    },
    sub() {
      let values = [this.value1, this.value2]

      this.calcResult = waSub(...values)
      this.operator = "-"
    },
    mul() {
      let values = [this.value1, this.value2]

      this.calcResult = waMulti(...values)
      this.operator = "*"
    },
    div() {
      let values = [this.value1, this.value2]

      this.calcResult = waDivi(...values)
      this.operator = "/"
    }
  }
}
</script>

<style scoped>
</style>
