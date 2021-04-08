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
      this.calcResult = this.$go.add(this.value1, this.value2)
      this.operator = "+"
    },
    sub() {
      this.calcResult = this.$go.sub(this.value1, this.value2)
      this.operator = "-"
    },
    mul() {
      this.calcResult = this.$go.multi(...[this.value1, this.value2])
      this.operator = "*"
    },
    div() {
      this.calcResult = this.$go.divi(...[this.value1, this.value2])
      this.operator = "/"
    }
  }
}
</script>

<style scoped>
</style>
