<template>
  <div>
    <nav class="navbar navbar-light bg-white">
      <div class="container">
        <nuxt-link class="navbar-brand font-weight-bold" to="/">
          Lamy
        </nuxt-link>
        <form class="d-flex">
          <nuxt-link
            v-if="!hasToken"
            class="btn btn-sm bg-transparent"
            to="/login"
          >
            Sign In
          </nuxt-link>
          <nuxt-link
            v-if="!hasToken"
            class="btn btn-sm bg-transparent"
            to="/register"
          >
            Sign Up
          </nuxt-link>
          <nuxt-link
            v-if="hasToken"
            class="btn btn-sm bg-transparent"
            to="/create"
          >
            Create
          </nuxt-link>
          <nuxt-link
            v-if="hasToken"
            class="btn btn-sm bg-transparent"
            to="/panel"
          >
            Panel
          </nuxt-link>
          <button
            v-if="hasToken"
            class="btn btn-sm btn-primary"
            @click="LogOut"
          >
            Sign Out
          </button>
        </form>
      </div>
    </nav>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions, mapGetters } from 'vuex'

export default Vue.extend({
  computed: {
    ...mapGetters({
      hasToken: 'hasToken',
    }),
  },
  methods: {
    ...mapActions({
      LogOut: 'LogOut',
    }),
  },
})
</script>
