<template>
  <div class="is-full-height d-flex justify-content-center align-items-center">
    <div class="container limit">
      <div class="row">
        <div class="col">
          <div class="card border-0 px-4 py-4 shadow">
            <h3 class="text-center text-muted">Create account</h3>
            <input
              v-model="localEmail"
              class="mt-3 form-control"
              type="email"
              placeholder="Enter your email"
              @input="handleRegisterEmail"
            />
            <input
              v-model="localPassword"
              class="mt-3 form-control"
              type="text"
              placeholder="Enter your password"
              @input="handleRegisterPassword"
            />
            <button class="btn btn-bg-blue mt-4" @click="SignUp">
              Sign Up
            </button>
            <hr class="dropdown-divider mt-3" />
            <nuxt-link to="/login" class="btn btn-bg-blue-darken mt-2"
              >Sign In</nuxt-link
            >
            <nuxt-link to="/" class="btn btn-light mt-3">Go back</nuxt-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapMutations, mapActions } from 'vuex'

export default Vue.extend({
  data: () => ({
    localEmail: '',
    localPassword: '',
  }),
  methods: {
    ...mapMutations({
      setRegisterEmail: 'SET_REGISTER_EMAIL',
      setRegisterPassword: 'SET_REGISTER_PASSWORD',
    }),
    ...mapActions({
      SignUp: 'SignUp',
    }),
    handleRegisterEmail() {
      this.setRegisterEmail(this.localEmail)
    },
    handleRegisterPassword() {
      this.setRegisterPassword(this.localPassword)
    },
  },
})
</script>

<style scoped>
.is-full-height {
  height: 100vh;
  width: 100vw;
  background-image: url('https://res.cloudinary.com/dj5iihhqv/image/upload/v1601675868/UPSPOT/4165818_br1jnr.jpg');
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  position: relative;
}
.limit {
  max-width: 400px;
}
.shadow {
  -webkit-box-shadow: 3px 3px 5px 6px #ccc; /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow: 3px 3px 5px 6px #ccc; /* Firefox 3.5 - 3.6 */
  box-shadow: 3px 3px 5px 6px #ccc; /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
}
.btn-bg-blue {
  background: #0540f2;
  color: white;
}
.btn-bg-blue-darken {
  background: #0f0140;
  color: white;
}
</style>
