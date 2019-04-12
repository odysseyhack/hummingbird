<template>
  <div id="app-home">
    <div class="logo">
      <img src="../assets/logo.png">
    </div>
    <div class="settings">
      <a v-on:click="goToSettings">
        <ons-icon style="color: whitesmoke;" size="25px" icon="fa-bars"></ons-icon>
      </a>
    </div>
  </div>
</template>

<script>
import ProgressBar from "vue-simple-progress";
export default {
  name: "Home",
  components: {
    ProgressBar
  },
  data() {
    return {
      cameras_online: [],
      enable_scan: false,
      enable_download: false,
      watchdog: null,
      clickCounter: 0,
      clickTimer: null
    };
  },
  computed: {
    getCamerasOnline: function() {
      return (this.cameras_online.length * 100) / 48;
    },
    getCamerasDownloaded: function() {
      return (this.$store.getters.getDownloadedCameras.length * 100) / 48;
    },
    cameras_downloaded: function() {
      return this.$store.getters.getDownloadedCameras;
    }
  },
  mounted: function() {
    if (this.$store.getters.getCameraStatus) {
      this.enable_scan = true;
    }
    if (this.$store.getters.getStatus.downloading) {
      this.enable_download = true;
    }
  },
  mqtt: {
    "health/#": function(data, topic) {
      try {
        let id = topic.split("/")[1];
        if (id != "main") {
          if (!this.$store.getters.getCameraStatus) {
            let id = topic.split("/")[1];
            if (this.cameras_online.indexOf(id) == -1) {
              this.cameras_online.push(id);
            }
            if (this.cameras_online.length > 46) {
              this.enable_scan = true;
              this.$store.dispatch("updateCameraStatus", true);
              this.$mqtt.publish("led/on", "255,255,255");
            }
          }
        }
      } catch (err) {
        console.log(err);
      }
    },
    "download/finished/#": function(data, topic) {
      try {
        if (this.$store.getters.getStatus.downloading) {
          if (this.$store.getters.getDownloadedCameras.length == 48) {
            this.enable_scan = true;
            this.enable_download = false;
            this.$store.dispatch("updateCameraStatus", true);
            this.$store.dispatch("updateDownloadStatus", false);
            this.$mqtt.publish("led/on", "255,255,255");
          }
        }
      } catch (err) {
        console.log(err);
      }
    },
    "download/finished/all": function(data, topic) {
      this.enable_scan = true;
      this.enable_download = false;
      this.$store.dispatch("updateCameraStatus", true);
      this.$store.dispatch("updateDownloadStatus", false);
    },
    "upload/ready/#": function(data, topic) {
      console.log("upload/ready triggered");
      this.enable_scan = true;
      this.enable_download = false;
      this.$store.dispatch("updateCameraStatus", true);
      this.$store.dispatch("updateDownloadStatus", false);
      this.$store.dispatch("resetCameraDownloaded", []);
    }
  },
  methods: {
    goToSettings: function() {
      console.log(" -- Go To Settings");
      this.$router.push("/settings");
    }
  }
};
</script>

<style>
.loader .vue-simple-progress,
.loader .vue-simple-progress-bar {
  border-radius: 5px !important;
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
  top: 120px;
  display: block;
  opacity: 0;
  -webkit-animation: fadein 0.7s forwards;
  animation: fadein 0.7s forwards;
  animation-delay: 0s;
}

@keyframes fadein {
  0% {
    top: 115px;
    opacity: 0;
  }
  100% {
    top: 120px;
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
  top: 20px;
  right: 30px;
  display: block;
  -webkit-animation: fadeinsimple 0.4s forwards;
  animation: fadeinsimple 0.4s forwards;
  animation-delay: 0s;
}

.settings img {
  width: 35px;
  height: auto;
}
</style>
