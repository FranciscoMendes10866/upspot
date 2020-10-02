<template>
  <div>
    <form class="row row-cols-lg-auto g-3 align-items-center">
      <div class="col-12">
        <input
          type="text"
          class="form-control"
          id="exampleFormControlInput1"
          style="width: 400px"
          placeholder="Encontre o seu próximo evento."
        />
      </div>

      <div class="col-12">
        <label class="visually-hidden" for="inlineFormSelectPref"
          >Preference</label
        >
        <select class="form-select" id="inlineFormSelectPref">
          <option selected>Coimbra</option>
          <option value="1">Lisboa</option>
          <option value="2">Porto</option>
        </select>
      </div>

      <div class="col-12">
        <button type="submit" class="btn btn-outline-primary">Procurar</button>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>
