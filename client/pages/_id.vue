<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar />
    </section>
    <!-- Header section -->
    <section class="py-5 border-top border-bottom border-1">
      <div class="container">
        <p class="text-muted">{{ event.date }}</p>
        <p class="h1 font-weight-bold mt-2 mb-2">{{ event.title }}</p>
        <small class="text-muted">Hosted by:</small>
        <a :href="event.hostURL" target="blank" class="link-primary dec"
          ><small>{{ event.hostName }}</small></a
        >
      </div>
    </section>
    <!-- Event info section -->
    <section class="bg-light py-5">
      <div class="container">
        <div class="d-flex justify-content-around flex-wrap">
          <!-- Left -->
          <div class="mt-3">
            <div class="img-w">
              <img class="rounded img-fluid" :src="event.img" alt />
              <p class="h5 mt-4 mb-4 font-weight-bold">Details</p>
              <p>
                <small>
                  It is a long established fact that a reader will be distracted
                  by the readable content of a page when looking at its layout.
                  The point of using Lorem Ipsum is that it has a more-or-less
                  normal distribution of letters, as opposed to using 'Content
                  here, content here', making it look like readable English.
                  Many desktop publishing packages and web page editors now use
                  Lorem Ipsum as their default model text, and a search for
                  'lorem ipsum' will uncover many web sites still in their
                  infancy. Various versions have evolved over the years,
                  sometimes by accident, sometimes on purpose (injected humour
                  and the like).
                </small>
              </p>
            </div>
          </div>
          <!-- Right -->
          <div class="mt-3">
            <div class="card border-0 shadow w px-3 py-3">
              <span class="text-muted mb-2">{{ event.type }}</span>
              <span>10:00 AM to 7:00 PM</span>
              <button
                v-if="attend === null"
                v-show="hasToken"
                class="btn btn-success mt-3"
                @click.prevent="Validate"
              >
                Validate
              </button>
              <button
                v-if="attend === false"
                class="btn btn-primary mt-3"
                @click.prevent="Attend"
              >
                Attend
              </button>
              <button
                v-if="attend === true"
                class="btn btn-danger mt-3"
                @click.prevent="NotAttend"
              >
                Not Attend
              </button>
              <a
                v-if="attend === true"
                :href="event.link"
                target="blank"
                class="btn btn-primary mt-2"
              >
                Visit
              </a>
              <nuxt-link to="/" class="btn btn-light mt-2">Go back</nuxt-link>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- Footer -->
    <section>
      <Footer class="border-top" />
    </section>
    {{ id }}
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'

import Navbar from '../components/Navbar.vue'
import Footer from '../components/Footer.vue'

export default Vue.extend({
  components: {
    Navbar,
    Footer,
  },
  data: () => ({
    event: {},
    attend: null,
  }),
  computed: {
    ...mapGetters({
      hasToken: 'hasToken',
    }),
  },
  created() {
    this.$axios
      .get(`http://localhost:3333/api/v1/events/${this.$route.params.id}`)
      .then((res: any) => {
        this.event = res.data
      })
  },
  methods: {
    async Validate() {
      await this.$axios.get(`http://localhost:3333/api/v1/attends/${this.$route.params.id}`, {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        })
        .then((res: any) => {
          const response = res.data.error === 'Empty array.'
          if (response === true) {
            this.attend = false
          } else {
            this.attend = true
          }
        })
    },
    async NotAttend() {
      await this.$axios
        .delete(
          `http://localhost:3333/api/v1/attends/${this.$route.params.id}`,
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.Token}`,
            },
          }
        )
        .then(() => {
          this.attend = false
        })
    },
    async Attend() {
      await this.$axios
        .post(`http://localhost:3333/api/v1/attends/${this.$route.params.id}`, {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        })
        .then(() => {
          this.attend = true
        })
    },
  },
})
</script>

<style scoped>
.shadow {
  -webkit-box-shadow: 2px 2px 2px 2px #ccc; /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow: 2px 2px 2px 2px #ccc; /* Firefox 3.5 - 3.6 */
  box-shadow: 2px 2px 2px 2px #ccc; /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
}
.w {
  width: 18rem;
}
.dec {
  text-decoration: none;
}
.img-w {
  max-width: 500px;
}
p {
  text-align: justify;
  text-justify: inter-word;
  margin: 0;
}
</style>
