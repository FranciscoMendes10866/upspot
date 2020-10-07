<template>
  <div>
    <form class="row row-cols-lg-auto g-3 align-items-center">
      <div class="col-12">
        <input
          type="text"
          class="form-control limit"
          id="exampleFormControlInput1"
          placeholder="Find your next event."
        />
      </div>

      <div class="col-12">
        <select class="form-select" id="inlineFormSelectPref">
          <option selected>Zurich</option>
          <option value="1">Luzern</option>
          <option value="2">Glarus</option>
        </select>
      </div>

      <div class="col-12">
        <button type="submit" class="btn btn-outline-primary">Search</button>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style scoped>
.limit {
  width: 400px;
}
@media only screen and (max-width: 600px) {
  .limit {
    width: 300px;
  }
}
</style>
