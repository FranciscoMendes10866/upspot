import { MutationTree, GetterTree, ActionTree } from 'vuex'

// Init states
export const state = () => ({
  // Register
  registerEmail: '',
  registerPassword: '',
  // Login
  loginEmail: '',
  loginPassword: '',
  // Data
  Token: '',
})

// Root State
export type RootState = ReturnType<typeof state>

// Mutations
export const mutations: MutationTree<RootState> = {
  // Register
  SET_REGISTER_EMAIL: (state, registerEmail: string) =>
    (state.registerEmail = registerEmail),
  SET_REGISTER_PASSWORD: (state, registerPassword: string) =>
    (state.registerPassword = registerPassword),
  // Login
  SET_LOGIN_EMAIL: (state, loginEmail: string) =>
    (state.loginEmail = loginEmail),
  SET_LOGIN_PASSWORD: (state, loginPassword: string) =>
    (state.loginPassword = loginPassword),
  // Data
  SET_TOKEN: (state, Token: string) => (state.Token = Token),
}

// Getters
export const getters: GetterTree<RootState, RootState> = {
  hasToken: (state) => !!state.Token,
}

// Actions
export const actions: ActionTree<RootState, RootState> = {
  async SignUp({ commit, state }) {
    const clientState = {
      email: state.registerEmail,
      password: state.registerPassword,
    }
    const send = await this.$axios.$post(
      'http://localhost:3333/api/v1/auth/sign_up',
      clientState
    )
    if (!send) {
      // eslint-disable-next-line no-console
      console.log('Register Error.')
    } else {
      commit('SET_REGISTER_EMAIL', '')
      commit('SET_REGISTER_PASSWORD', '')
      this.$router.push('/login')
    }
  },
  async SignIn({ commit, state }) {
    const clientState = {
      email: state.loginEmail,
      password: state.loginPassword,
    }
    const send = await this.$axios.$post(
      'http://localhost:3333/api/v1/auth/sign_in',
      clientState
    )
    if (!send) {
      // eslint-disable-next-line no-console
      console.log('Login Error.')
    } else {
      commit('SET_LOGIN_EMAIL', '')
      commit('SET_LOGIN_PASSWORD', '')
      commit('SET_TOKEN', send.token)
      this.$router.push('/')
    }
  },
  LogOut({ commit }) {
    commit('SET_TOKEN', '')
    this.$router.push('/')
  },
}
