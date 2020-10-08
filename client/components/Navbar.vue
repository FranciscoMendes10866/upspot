<template>
  <div>
    <nav class="navbar navbar-light bg-white">
      <div class="container">
        <nuxt-link class="navbar-brand font-weight-bold" to="/">
          Lamy
        </nuxt-link>
        <form class="d-flex">
          <nuxt-link class="btn btn-sm bg-transparent" to="/login">
            Sign In
          </nuxt-link>
          <nuxt-link class="btn btn-sm bg-transparent" to="/register">
            Sign Up
          </nuxt-link>
          <nuxt-link class="btn btn-sm bg-transparent" to="/panel">
            Panel
          </nuxt-link>
          <button class="btn btn-sm btn-primary">Sign Out</button>
        </form>
      </div>
    </nav>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>
