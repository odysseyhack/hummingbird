<template>
  <div id="app-home">
    <!-- <div class="logo">
      <img src="../assets/logo.png">
    </div>-->
    <div class="settings">
      <span class="title">Hummingbird</span>
      <a v-on:click="goToSettings">
        <img src="../assets/settings.svg">
      </a>
      <a v-on:click="logOut">
        <img src="../assets/logout.svg">
      </a>
    </div>
    <div class="sidebar">
      <template v-for="(menu,key,index) in menus">
        <div
          @click="changeFilter"
          class="sidebar-item"
          v-bind:key="index"
          v-bind:class="{ sidebarActive: menu.active }"
        >
          <img class="sidebar-icon" v-bind:src="menu.icon">
          <span class="sidebar-label">{{menu.label}}</span>
        </div>
      </template>
    </div>
    <div class="map">
      <GmapMap
        :center="this.center"
        :zoom="this.zoom"
        map-type-id="roadmap"
        style="width: 100%; height: 100%"
        :options="{
          zoomControl: true,
          mapTypeControl: true,
          scaleControl: false,
          streetViewControl: false,
          rotateControl: false,
          fullscreenControl: false
        }"
      >
        <GmapMarker
          :key="index"
          v-for="(m, index) in markers"
          :position="m.position"
          :clickable="true"
          :draggable="false"
          @click="center=m.position"
        />
        <template v-for="(marker,key,index) in markers">
          <gmap-custom-marker :marker="marker" v-bind:key="index">
            <img src="../assets/icons/Water.svg">
          </gmap-custom-marker>
        </template>
      </GmapMap>
    </div>
  </div>
</template>

<script>
import GmapCustomMarker from "vue2-gmap-custom-marker";

export default {
  name: "Home",
  components: {
    "gmap-custom-marker": GmapCustomMarker
  },
  data() {
    return {
      zoom: 10,
      center: { lat: 53.269692, lng: 6.818274 },
      markers: [
        { lat: 53.289486, lng: 6.838749 },
        { lat: 53.218609, lng: 6.550166 }
      ],
      menus: [
        {
          icon: require("../assets/icons/Water icon.svg"),
          status: true,
          label: "test",
          active: false
        },
        {
          icon: require("../assets/icons/Electricity.svg"),
          status: true,
          label: "test",
          active: true
        },
        {
          icon: require("../assets/icons/Fire icon.svg"),
          status: true,
          label: "test",
          active: false
        },
        {
          icon: require("../assets/icons/First aid icon.svg"),
          status: true,
          label: "test",
          active: false
        }
      ]
    };
  },
  computed: {},
  mounted: function() {
    window.setTimeout(() => {
      let coord = { lat: 53.23311, lng: 6.586164 };
      this.markers.push(coord);
      this.zoom = 11;
      this.center = coord;
    }, 1000);
    window.setTimeout(() => {
      let coord = { lat: 53.242514, lng: 6.65732 };
      this.markers.push(coord);
      this.zoom = 11;
      this.center = coord;
    }, 5000);
  },
  mqtt: {},
  methods: {
    changeFilter: function(msg) {
      console.log(msg);
    },
    goToSettings: function() {
      console.log(" -- Go To Settings");
      this.$router.push("/settings");
    },
    logOut: function() {
      console.log(" -- Log out?");
      this.$ons.notification
        .confirm({
          title: "Confirm",
          message: "Are you sure you want to exit Hummingbird?"
        })
        .then(arg => {
          if (arg == 1) {
            // yes
            console.log(" -- Shutting down");
          }
        });
    }
  }
};
</script>

<style>
.loader .vue-simple-progress,
.loader .vue-simple-progress-bar {
  border-radius: 5px !important;
}
@font-face {
  font-family: "Skycoin";
  font-style: normal;
  font-weight: 700;
  src: url("../assets/fonts/skycoin/Skycoin-Bold.woff2") format("woff2"),
    url("../assets/fonts/skycoin/Skycoin-Bold.woff") format("woff");
}
</style>

<style scoped>
.loader .vue-simple-progress,
.loader .vue-simple-progress-bar {
  border-radius: 5px !important;
}

#app-home {
  position: relative;
  top: 0;
  text-align: center;
}

.sidebar {
  display: block;
  position: absolute;
  float: left;
  top: 50px;
  width: 120px;
  height: calc(100vh-50px);
  background: -webkit-linear-gradient(bottom, black, rgb(73, 144, 168));
}

.sidebar-icon {
  height: 30px;
  margin: 0 auto;
  top: 20px;
  position: relative;
  display: block;
}

.sidebar-label {
  position: relative;
  display: block;
  top: 25px;
  font-size: 12px;
  color: whitesmoke;
}

.sidebar-item {
  opacity: 0.5;
  position: relative;
  display: block;
  width: 100%;
  height: 100px;
  background: -webkit-linear-gradient(
    bottom,
    rgb(20, 76, 144, 0.7),
    rgb(73, 144, 168)
  );
  -webkit-transition: 0.1s all ease-in-out;
}

.sidebar-item:hover {
  opacity: 1;
}

.sidebarActive {
  background: -webkit-linear-gradient(
    0,
    rgba(0, 114, 255, 0.7),
    rgba(0, 195, 255, 1)
  );
  opacity: 1;
}

.map {
  display: block;
  position: absolute;
  right: 0;
  top: 50px;
  width: calc(100vw - 120px);
  height: calc(100vh - 50px);
  background-color: black;
}

.loader {
  position: relative;
  top: 170px;
  width: 190px;
  margin: 0 auto;
}

.loader-status {
  position: relative;
  top: 160px;
  font-weight: bold;
  color: white;
}

.menu {
  position: relative;
  top: 120px;
  display: block;
}

.logo {
  position: relative;
  float: left;
  display: block;
  opacity: 0;
  left: 20px;
  margin-top: 10px;
  -webkit-animation: fadein 0.7s forwards;
  animation: fadein 0.7s forwards;
  animation-delay: 0s;
}

.logo img {
  width: 30px;
}

@keyframes fadein {
  0% {
    top: 0;
    opacity: 0;
  }
  100% {
    top: 0;
    opacity: 1;
  }
}

@keyframes fadeinsimple {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

.settings {
  position: absolute;
  height: 50px;
  line-height: 50px;
  right: 30px;
  display: block;
  -webkit-animation: fadeinsimple 0.4s forwards;
  animation: fadeinsimple 0.4s forwards;
  animation-delay: 0s;
  float: left;
}

.settings img {
  position: relative;
  width: 20px;
  top: 3px;
  height: auto;
  padding-left: 20px;
}

.settings .title {
  position: relative;
  /* top: -2px; */
  font-family: "Skycoin";
  font-size: 20px;
  color: whitesmoke;
  padding-right: 20px;
}
</style>
