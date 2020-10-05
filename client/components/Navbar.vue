<template>
  <div>
    <nav class="navbar navbar-light bg-white fixed-top">
      <div class="container">
        <a class="navbar-brand font-weight-bold">Lamy</a>
        <form class="d-flex">
          <nuxt-link class="btn bg-transparent" to="/login">
            Iniciar sessão
          </nuxt-link>
          <nuxt-link class="btn bg-transparent" to="/register">
            Criar conta
          </nuxt-link>
        </form>
      </div>
    </nav>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>
