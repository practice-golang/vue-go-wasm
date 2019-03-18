<template>
  <div class="hello">
    <h1>WASM Arithmatic ^-^</h1>

    <input type="text" v-model="value1"> <span>{{operator}}</span>
    <input type="text" v-model="value2">
    =
    <input type="text" v-model="calcResult" disabled>
    <hr>
    <button @click="add();" id="addButton">Add</button>
    <button @click="sub();" id="subButton">Sub</button>
    <button @click="mul();" id="subButton">Mult</button>
    <button @click="div();" id="subButton">Div</button>
  </div>
</template>

<script>
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
  methods: {
    async add() {
      let values = [this.value1, this.value2]

      // eslint-disable-next-line
      // this.result = await this.$go.add(...values)
      // eslint-disable-next-line
      this.calcResult = await add(...values)

      this.operator = "+"
    },
    async sub() {
      let values = [this.value1, this.value2]

      // eslint-disable-next-line
      this.calcResult = await sub(...values)

      this.operator = "-"
    },
    async mul() {
      let values = [this.value1, this.value2]

      // eslint-disable-next-line
      this.calcResult = await multi(...values)

      this.operator = "*"
    },
    async div() {
      let values = [this.value1, this.value2]

      // eslint-disable-next-line
      this.calcResult = await divi(...values)

      this.operator = "/"
    }
  }
}
</script>

<style scoped>
</style>
