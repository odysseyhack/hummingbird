<template>
  <div id="body-background">
    <div id="app">
      <transition name="fade" mode="out-in">
        <router-view class="view"/>
      </transition>
    </div>
  </div>
</template>

<script>
import offline from "v-offline";

export default {
  name: "App",
  components: {
    offline
  },
  data() {
    return {
      counter: 0,
      imagePath: "",
      shutdownTimer: null,
      images: [],
      toastVisible: false,
      network: true
    };
  },
  created: function() {
    this.subscribe();
  },
  mqtt: {
    "hiber/faker": function(data, topic) {
      console.log(JSON.parse(data));
    }
  },
  methods: {
    handleConnectivityChange: function(status) {
      this.$ons.notification.toast("Network changed: " + status, {
        timeout: 1000,
        animation: "fall"
      });
      this.toastVisible = true;
      this.network = status;
      let insideThis = this;
      window.setTimeout(function() {
        insideThis.toastVisible = false;
      }, 4000);
    },
    parseByteArray: function(data) {
      return String.fromCharCode.apply(null, data);
    },
    show: function() {
      const viewer = this.$el.querySelector(".images").$viewer;
      viewer.show();
    },
    update: function() {
      const viewer = this.$el.querySelector(".images").$viewer;
      viewer.update();
      viewer.show();
      let a = this;
    },
    subscribe: function() {
      this.$mqtt.subscribe("hiber/faker", {});
    }
  }
};
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0s ease-in-out;
}

.fade-enter,
.fade-leave-active {
  opacity: 0.5;
}

.child-view {
  position: absolute;
  transition: all 0.3s cubic-bezier(0.55, 0, 0.1, 1);
}

.slide-left-enter,
.slide-right-leave-active {
  opacity: 0;
  -webkit-transform: translate(30px, 0);
  transform: translate(30px, 0);
}

.slide-left-leave-active,
.slide-right-enter {
  opacity: 0;
  -webkit-transform: translate(-30px, 0);
  transform: translate(-30px, 0);
}

html,
body {
  height: 100%;
  margin: 0;
}

.mano-button {
  font-family: "Open Sans", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  display: block;
  opacity: 1;
  width: 175px;
  background-color: rgba(255, 255, 255, 1) !important;
  border-radius: 500px;
  line-height: 100%;
  margin: 0;
  padding: 20px;
  text-align: center;
  height: 35px;
  margin-left: 0;
  font-size: 30px;
  color: black !important;
  transition: all 0.2s ease-in-out;
  text-decoration: none !important;
  -webkit-touch-callout: none;
  /* iOS Safari */
  -webkit-user-select: none;
  /* Safari */
  -khtml-user-select: none;
  /* Konqueror HTML */
  -moz-user-select: none;
  /* Firefox */
  -ms-user-select: none;
  /* Internet Explorer/Edge */
  user-select: none;
}

.mano-button:hover {
  background-color: rgba(255, 255, 255, 0.9);
  transition: all 0.2s ease-in-out;
  text-decoration: none !important;
}

body {
  width: 100%;
  min-height: 100%;
  background-color: black;
  top: 0;
  left: 0;
  position: relative;
  display: block;
  margin: 0;
  padding: 0;
  overflow: hidden;
}

#body-background {
  background-color: rgba(30, 34, 39, 1);
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  position: absolute;
  display: block;
  transition: all 0.13s ease-in-out;
  opacity: 1;
}

#background-wrapper {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  background: -webkit-radial-gradient(
    60% 70%,
    rgba(0, 114, 255, 0.7) 0%,
    rgba(0, 195, 255, 1) 90%
  );
  overflow: hidden;
  transition: all 1s ease-in-out;
  opacity: 1;
}

#app {
  font-family: "Open Sans", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: absolute;
  min-width: 100%;
  min-height: 100%;
  height: 100%;
  top: 0;
  left: 0;
  display: block;
  overflow: hidden;
}

/* #app img {
    opacity: 0;
    display: flex;
    align-items: center;
    padding: 5px;
    position: relative;
    margin: 10px auto;
    -webkit-animation: fadein 0.7s forwards;
    animation: fadein 0.7s forwards;
    animation-delay: 0.3s;
    -webkit-animation-delay: 0.3s;
      -khtml-user-select: none;
      -o-user-select: none;
      -moz-user-select: none;
      -webkit-user-select: none;
      user-select: none;
  } */

@keyframes fadein {
  0% {
    opacity: 0;
    top: -10px;
  }
  100% {
    opacity: 1;
    top: 0;
  }
}

/*!
   * Viewer.js v1.3.0
   * https://fengyuanchen.github.io/viewerjs
   *
   * Copyright 2015-present Chen Fengyuan
   * Released under the MIT license
   *
   * Date: 2018-10-25T12:41:52.459Z
   */

.viewer-zoom-in::before,
.viewer-zoom-out::before,
.viewer-one-to-one::before,
.viewer-reset::before,
.viewer-prev::before,
.viewer-play::before,
.viewer-next::before,
.viewer-rotate-left::before,
.viewer-rotate-right::before,
.viewer-flip-horizontal::before,
.viewer-flip-vertical::before,
.viewer-fullscreen::before,
.viewer-fullscreen-exit::before,
.viewer-close::before {
  background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAARgAAAAUCAYAAABWOyJDAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAALEgAACxIB0t1+/AAAABx0RVh0U29mdHdhcmUAQWRvYmUgRmlyZXdvcmtzIENTNui8sowAAAQPSURBVHic7Zs/iFxVFMa/0U2UaJGksUgnIVhYxVhpjDbZCBmLdAYECxsRFBTUamcXUiSNncgKQbSxsxH8gzAP3FU2jY0kKKJNiiiIghFlccnP4p3nPCdv3p9778vsLOcHB2bfveeb7955c3jvvNkBIMdxnD64a94GHMfZu3iBcRynN7zAOI7TG15gHCeeNUkr8zaxG2lbYDYsdgMbktBsP03jdQwljSXdtBhLOmtjowC9Mg9L+knSlcD8TNKpSA9lBpK2JF2VdDSR5n5J64m0qli399hNFMUlpshQii5jbXTbHGviB0nLNeNDSd9VO4A2UdB2fp+x0eCnaXxWXGA2X0au/3HgN9P4LFCjIANOJdrLr0zzZ+BEpNYDwKbpnQMeAw4m8HjQtM6Z9qa917zPQwFr3M5KgA6J5rTJCdFZJj9/lyvGhsDvwFNVuV2MhhjrK6b9bFiE+j1r87eBl4HDwCF7/U/k+ofAX5b/EXBv5JoLMuILzf3Ap6Z3EzgdqHMCuF7hcQf4HDgeoHnccncqdK/TvSDWffFXI/exICY/xZyqc6XLWF1UFZna4gJ7q8BsRvgd2/xXpo6P+D9dfT7PpECtA3cnWPM0GXGFZh/wgWltA+cDNC7X+AP4GzjZQe+k5dRxuYPeiuXU7e1qwLpDz7dFjXKRaSwuMLvAlG8zZlG+YmiK1HoFqT7wP2z+4Q45TfEGcMt01xLoNZEBTwRqD4BLpnMLeC1A41UmVxsXgXeBayV/Wx20rpTyrpnWRft7p6O/FdqzGrDukPNtkaMoMo3FBdBSQMOnYBCReyf05s126fU9ytfX98+mY54Kxnp7S9K3kj6U9KYdG0h6UdLbkh7poFXMfUnSOyVvL0h6VtIXHbS6nOP+s/Zm9mvyXW1uuC9ohZ72E9uDmXWLJOB1GxsH+DxPftsB8B6wlGDN02TAkxG6+4D3TWsbeC5CS8CDFce+AW500LhhOW2020TRjK3b21HEmgti9m0RonxbdMZeVzV+/4tF3cBpP7E9mKHNL5q8h5g0eYsCMQz0epq8gQrwMXAgcs0FGXGFRcB9wCemF9PkbYqM/Bas7fxLwNeJPdTdpo4itQti8lPMqTpXuozVRVXPpbHI3KkNTB1NfkL81j2mvhDp91HgV9MKuRIqrykj3WPq4rHyL+axj8/qGPmTqi6F9YDlHOvJU6oYcTsh/TYSzWmTE6JT19CtLTJt32D6CmHe0eQn1O8z5AXgT4sx4Vcu0/EQecMydB8z0hUWkTd2t4CrwNEePqMBcAR4mrBbwyXLPWJa8zrXmmLEhNBmfpkuY2102xxrih+pb+ieAb6vGhuA97UcJ5KR8gZ77K+99xxeYBzH6Q3/Z0fHcXrDC4zjOL3hBcZxnN74F+zlvXFWXF9PAAAAAElFTkSuQmCC");
  background-repeat: no-repeat;
  color: transparent;
  display: block;
  font-size: 0;
  height: 20px;
  line-height: 0;
  width: 20px;
}

.viewer-zoom-in::before {
  background-position: 0 0;
  content: "Zoom In";
}

.viewer-zoom-out::before {
  background-position: -20px 0;
  content: "Zoom Out";
}

.viewer-one-to-one::before {
  background-position: -40px 0;
  content: "One to One";
}

.viewer-reset::before {
  background-position: -60px 0;
  content: "Reset";
}

.viewer-prev::before {
  background-position: -80px 0;
  content: "Previous";
}

.viewer-play::before {
  background-position: -100px 0;
  content: "Play";
}

.viewer-next::before {
  background-position: -120px 0;
  content: "Next";
}

.viewer-rotate-left::before {
  background-position: -140px 0;
  content: "Rotate Left";
}

.viewer-rotate-right::before {
  background-position: -160px 0;
  content: "Rotate Right";
}

.viewer-flip-horizontal::before {
  background-position: -180px 0;
  content: "Flip Horizontal";
}

.viewer-flip-vertical::before {
  background-position: -200px 0;
  content: "Flip Vertical";
}

.viewer-fullscreen::before {
  background-position: -220px 0;
  content: "Enter Full Screen";
}

.viewer-fullscreen-exit::before {
  background-position: -240px 0;
  content: "Exit Full Screen";
}

.viewer-close::before {
  background-position: -260px 0;
  content: "Close";
}

.viewer-container {
  bottom: 0;
  direction: ltr;
  font-size: 0;
  left: 0;
  line-height: 0;
  overflow: hidden;
  position: absolute;
  right: 0;
  -webkit-tap-highlight-color: transparent;
  top: 0;
  -webkit-touch-callout: none;
  -ms-touch-action: none;
  touch-action: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.viewer-container::-moz-selection,
.viewer-container *::-moz-selection {
  background-color: transparent;
}

.viewer-container::selection,
.viewer-container *::selection {
  background-color: transparent;
}

.viewer-container img {
  display: block;
  height: auto;
  max-height: none !important;
  max-width: none !important;
  min-height: 0 !important;
  min-width: 0 !important;
  width: 100%;
}

.viewer-canvas {
  bottom: 0;
  left: 0;
  overflow: hidden;
  position: absolute;
  right: 0;
  top: 0;
}

.viewer-canvas > img {
  height: auto;
  margin: 15px auto;
  max-width: 90% !important;
  width: auto;
}

.viewer-footer {
  bottom: 0;
  left: 0;
  overflow: hidden;
  position: absolute;
  right: 0;
  text-align: center;
}

.viewer-navbar {
  background-color: rgba(0, 0, 0, 0.5);
  overflow: hidden;
}

.viewer-list {
  -webkit-box-sizing: content-box;
  box-sizing: content-box;
  height: 50px;
  margin: 0;
  overflow: hidden;
  padding: 1px 0;
}

.viewer-list > li {
  color: transparent;
  cursor: pointer;
  float: left;
  font-size: 0;
  height: 50px;
  line-height: 0;
  opacity: 0.5;
  overflow: hidden;
  -webkit-transition: opacity 0.15s;
  transition: opacity 0.15s;
  width: 30px;
}

.viewer-list > li:hover {
  opacity: 0.75;
}

.viewer-list > li + li {
  margin-left: 1px;
}

.viewer-list > .viewer-loading {
  position: relative;
}

.viewer-list > .viewer-loading::after {
  border-width: 2px;
  height: 20px;
  margin-left: -10px;
  margin-top: -10px;
  width: 20px;
}

.viewer-list > .viewer-active,
.viewer-list > .viewer-active:hover {
  opacity: 1;
}

.viewer-player {
  background-color: #000;
  bottom: 0;
  cursor: none;
  display: none;
  left: 0;
  position: absolute;
  right: 0;
  top: 0;
}

.viewer-player > img {
  left: 0;
  position: absolute;
  top: 0;
}

.viewer-toolbar > ul {
  display: inline-block;
  margin: 0 auto 5px;
  overflow: hidden;
  padding: 3px 0;
}

.viewer-toolbar > ul > li {
  background-color: rgba(0, 0, 0, 0.5);
  border-radius: 50%;
  cursor: pointer;
  float: left;
  height: 24px;
  overflow: hidden;
  -webkit-transition: background-color 0.15s;
  transition: background-color 0.15s;
  width: 24px;
}

.viewer-toolbar > ul > li:hover {
  background-color: rgba(0, 0, 0, 0.8);
}

.viewer-toolbar > ul > li::before {
  margin: 2px;
}

.viewer-toolbar > ul > li + li {
  margin-left: 1px;
}

.viewer-toolbar > ul > .viewer-small {
  height: 18px;
  margin-bottom: 3px;
  margin-top: 3px;
  width: 18px;
}

.viewer-toolbar > ul > .viewer-small::before {
  margin: -1px;
}

.viewer-toolbar > ul > .viewer-large {
  height: 30px;
  margin-bottom: -3px;
  margin-top: -3px;
  width: 30px;
}

.viewer-toolbar > ul > .viewer-large::before {
  margin: 5px;
}

.viewer-tooltip {
  background-color: rgba(0, 0, 0, 0.8);
  border-radius: 10px;
  color: #fff;
  display: none;
  font-size: 12px;
  height: 20px;
  left: 50%;
  line-height: 20px;
  margin-left: -25px;
  margin-top: -10px;
  position: absolute;
  text-align: center;
  top: 50%;
  width: 50px;
}

.viewer-title {
  color: #ccc;
  display: inline-block;
  font-size: 12px;
  line-height: 1;
  margin: 0 5% 5px;
  max-width: 90%;
  opacity: 0.8;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-transition: opacity 0.15s;
  transition: opacity 0.15s;
  white-space: nowrap;
}

.viewer-title:hover {
  opacity: 1;
}

.viewer-button {
  background-color: rgba(0, 0, 0, 0.5);
  border-radius: 50%;
  cursor: pointer;
  height: 80px;
  overflow: hidden;
  position: absolute;
  right: -40px;
  top: -40px;
  -webkit-transition: background-color 0.15s;
  transition: background-color 0.15s;
  width: 80px;
}

.viewer-button:focus,
.viewer-button:hover {
  background-color: rgba(0, 0, 0, 0.8);
}

.viewer-button::before {
  bottom: 15px;
  left: 15px;
  position: absolute;
}

.viewer-fixed {
  position: fixed;
}

.viewer-open {
  overflow: hidden;
}

.viewer-show {
  display: block;
}

.viewer-hide {
  display: none;
}

.viewer-backdrop {
  background-color: rgba(0, 0, 0, 0.5);
}

.viewer-invisible {
  visibility: hidden;
}

.viewer-move {
  cursor: move;
  cursor: -webkit-grab;
  cursor: grab;
}

.viewer-fade {
  opacity: 0;
}

.viewer-in {
  opacity: 1;
}

.viewer-transition {
  -webkit-transition: all 0.3s;
  transition: all 0.3s;
}

@-webkit-keyframes viewer-spinner {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}

@keyframes viewer-spinner {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}

.viewer-loading::after {
  -webkit-animation: viewer-spinner 1s linear infinite;
  animation: viewer-spinner 1s linear infinite;
  border: 4px solid rgba(255, 255, 255, 0.1);
  border-left-color: rgba(255, 255, 255, 0.5);
  border-radius: 50%;
  content: "";
  display: inline-block;
  height: 40px;
  left: 50%;
  margin-left: -20px;
  margin-top: -20px;
  position: absolute;
  top: 50%;
  width: 40px;
  z-index: 1;
}

@media (max-width: 767px) {
  .viewer-hide-xs-down {
    display: none;
  }
}

@media (max-width: 991px) {
  .viewer-hide-sm-down {
    display: none;
  }
}

@media (max-width: 1199px) {
  .viewer-hide-md-down {
    display: none;
  }
}

/* Make the element pulse (grow large and small slowly) */

/* Usage
      .myElement {
          animation: pulsate 1s ease-out;
          animation-iteration-count: infinite;
          opacity: 1;
      }
  */

@-webkit-keyframes pulsate {
  0% {
    -webkit-transform: scale(0.1, 0.1);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    -webkit-transform: scale(1.2, 1.2);
    opacity: 0;
  }
}

/* Make the element's opacity pulse*/

/* Usage
      .myElement {
          animation: opacityPulse 1s ease-out;
          animation-iteration-count: infinite;
          opacity: 0;
      }
  */

@-webkit-keyframes opacityPulse {
  0% {
    opacity: 0.6;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0.6;
  }
}

/* Make the element's background pulse. I call this alertPulse because it is red. You can call it something more generic. */

/* Usage
      .myElement {
          animation: alertPulse 1s ease-out;
          animation-iteration-count: infinite;
          opacity: 1;
      }
  */

@-webkit-keyframes alertPulse {
  0% {
    background-color: #9a2727;
    opacity: 1;
  }
  50% {
    opacity: red;
    opacity: 0.75;
  }
  100% {
    opacity: #9a2727;
    opacity: 1;
  }
}

/* Make the element rotate infinitely. */

/*
  Usage
      .myElement {
          animation: rotating 3s linear infinite;
      }
  */

@keyframes rotating {
  from {
    -ms-transform: rotate(0deg);
    -moz-transform: rotate(0deg);
    -webkit-transform: rotate(0deg);
    -o-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  to {
    -ms-transform: rotate(360deg);
    -moz-transform: rotate(360deg);
    -webkit-transform: rotate(360deg);
    -o-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
</style>
