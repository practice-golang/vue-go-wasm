<template>
  <div class="hello">
    <h1>Arithmatic ^-^</h1>

    <input type="text" id="value1">
    
    <button onClick="add('value1', 'value2', 'result');" id="addButton">Add</button>
    <button onClick="subtract('value1', 'value2', 'result');" id="subtractButton">Subtract</button>
    
    <input type="text" id="value2">
    =
    <input type="text" id="result" disabled>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
