<template>
  <div class="SettingsPage">
    <div class="menu-holder">
      <v-ons-dialog @deviceBackButton="$event.callParentHandler()"/>
    </div>
    <div class="footer-info">
      <a>v{{ getVersion }}</a>
    </div>
    <router-link to="/" class="back-butt">
      <ons-icon style="color: whitesmoke;" size="35px" icon="md-arrow-back"/>
    </router-link>
  </div>
</template>

<script>
export default {
  name: "SettingsPage",
  components: {},
  computed: {
    getVersion() {
      return this.$store.getters.getVersion;
    }
  },
  data() {
    return {};
  },
  methods: {
    goBack: function() {
      this.$router.push("/");
    }
  }
};
</script>

<style scoped>
.SettingsPage {
  top: 0;
  left: 0;
  height: 100%;
  text-align: center;
  background-color: rgba(255, 255, 255, 0);
}

.back-butt {
  position: absolute;
  top: 15px;
  left: -20px;
  width: 120px;
  color: black !important;
}

.menu-holder {
  position: absolute;
  top: 130px;
  left: 50%;
  display: block;
  width: 95%;
  height: 200px;
  margin-left: -47.5%;
  background-color: transparent;
}

.menu-holder > div {
  padding: 20px;
}

.button-restart,
.button-shutdown,
.button-debug {
  position: relative;
  display: block;
  float: left;
  width: 28%;
  padding: 0;
  margin: 0;
  text-align: center;
  background-color: transparent;
}

.button-restart span,
.button-shutdown span,
.button-debug span {
  position: relative;
  top: 20px;
  display: block;
  font-size: 30px;
  color: whitesmoke;
}

.button-debug {
  opacity: 1;
}

.button-debug img {
  transform: rotate(180deg);
}

.footer-info {
  display: block;
  position: absolute;
  left: 10px;
  width: 150px;
  text-align: left;
  padding: 0 10px 0 0px;
  height: 25px;
  bottom: 0;
  opacity: 1;
  /* -webkit-animation: fadein 0.7s forwards;
    animation: fadein 0.3s forwards;
    animation-delay: 0s; */
  font-size: 11px;
}

.footer-info a {
  text-decoration: none;
  color: whitesmoke !important;
}

@keyframes fadein {
  0% {
    opacity: 0;
    bottom: -10px;
  }
  100% {
    opacity: 1;
    bottom: 0;
  }
}
</style>