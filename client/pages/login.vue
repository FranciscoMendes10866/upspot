<template>
  <div class="is-full-height">
    <div class="container limit">
      <div class="row">
        <div class="col">
          <div class="card px-4 py-4">
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
            <button class="btn btn-dark mt-2">Criar</button>
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
  display: flex;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  align-items: center;
}
.limit {
  max-width: 400px;
}
</style>
