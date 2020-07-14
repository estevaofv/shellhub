<template>
  <v-form>
    <v-container>
      <v-row
        align="center"
        justify="center"
        class="mt-4"
      >
        <v-col
          sm="8"
        >
          <div
            class="mt-6 pl-4 pr-4"
          >
            <v-row>
              <v-col md="auto">
                <v-card
                  tile
                  :elevation="0"
                >
                  Tenant ID:
                </v-card>
              </v-col>
              <v-col
                md="auto"
                class="ml-auto"
              >
                <v-card
                  class="auto"
                  tile
                  :elevation="0"
                >
                  <v-chip>
                    <span>
                      {{ tenant }}
                    </span>
                    <v-icon
                      v-clipboard="tenant"
                      v-clipboard:success="() => { copySnack = true; }"
                      right
                    >
                      mdi-content-copy
                    </v-icon>
                  </v-chip>
                </v-card>
              </v-col>
            </v-row>
          </div>

          <v-divider />

          <div
            class="mt-6 pl-4 pr-4"
          >
            <v-text-field
              v-model="currentUser"
              label="Username"
              type="text"
              readonly
              class="mt-6"
            />

            <v-text-field
              v-model="email"
              label="E-mail"
              type="text"
              readonly
            />

            <v-text-field
              v-model="password"
              :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
              :type="show1 ? 'text' : 'password'"
              label="Password"
              readonly
              disabled
              @click:append="show1 = !show1"
            />
            <v-btn
              @click="postUser()"
            >
              Greet
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>

export default {
  name: 'SettingProfile',

  data() {
    return {
      show1: false,
      password: 'XXXXXXXXXX',
    };
  },

  computed: {
    currentUser() {
      return this.$store.getters['auth/currentUser'];
    },

    tenant() {
      return this.$store.getters['auth/tenant'];
    },

    email() {
      return this.$store.getters['auth/email'];
    },
  },

  methods: {
    postUser() {
      this.$store.dispatch('users/setUser', 'teste');
    },
  },

};
</script>
