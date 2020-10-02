<template>
  <div class="is-full-height d-flex justify-content-center align-items-center">
    <div class="container limit">
      <div class="row">
        <div class="col">
          <div class="card px-4 py-4 shadow">
            <h1 class="text-center mb-4 text-muted">Iniciar sessão</h1>
            <input
              class="form-control"
              type="email"
              placeholder="Insira o seu email"
            />
            <input
              class="mt-3 form-control"
              type="text"
              placeholder="Insira a sua palavra passe"
            />
            <button class="btn btn-primary mt-4">Iniciar</button>
            <hr class="dropdown-divider mt-3" />
            <nuxt-link to="/register" class="btn btn-dark mt-2"
              >Criar</nuxt-link
            >
            <nuxt-link to="/" class="btn btn-light mt-3">Voltar</nuxt-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({})
</script>

<style scoped>
.is-full-height {
  height: 100vh;
  width: 100vw;
}
.limit {
  max-width: 400px;
}
.shadow {
  -webkit-box-shadow: 3px 3px 5px 6px #ccc; /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow: 3px 3px 5px 6px #ccc; /* Firefox 3.5 - 3.6 */
  box-shadow: 3px 3px 5px 6px #ccc; /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
}
</style>
