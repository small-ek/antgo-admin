<template>
  <div>
    <h3>语音通话</h3>
    <audio ref="localAudio" style="width: 400px;height: 300px;" @click="toggleFullscreen"></audio>
    <div style="text-align: left">
      自己ID<input type="text" v-model="myPeerid"/>
      <br>
      对方ID<input type="text" v-model="youPeerid"/>
      <br>
      <button @click="startCaller">语音通话</button>
      <button @click="startCall">确认</button>
      <button @click="endCall">挂断</button>
    </div>
    <audio ref="remoteAudio" style="width: 400px;height: 300px;" @click="toggleFullscreen"></audio>
  </div>
</template>

<script>
import {Peer} from 'peerjs'

export default {
  data() {
    return {
      myPeerid: '',
      youPeerid: '',
      localStream: null,
      remoteStream: null,
      peer: null,
      call: null
    }
  },
  mounted() {
    this.start()
  },

  methods: {
    start() {
      this.getUserMedia({audio: true}, (stream) => {
        this.localStream = stream
        console.log(stream)
        this.$refs.localAudio.srcObject = stream
        // this.$refs.localAudio.autoplay = true
        // this.$refs.localAudio.play()
      }, (err) => {
        console.log('获取本地流失败', err)
      })

      this.peer = new Peer()
      this.peer.on('open', (id) => {
        this.myPeerid = id
      })
    },
    startCaller() {


      this.peer.on('call', (call) => {
        console.log(call)
        this.call = call
        call.answer(this.localStream)
        call.on('stream', (remoteStream) => {
          this.remoteStream = remoteStream
          this.$refs.remoteAudio.srcObject = remoteStream
          this.$refs.remoteAudio.autoplay = true
        })
      })
    },

    toggleFullscreen() {
      const audio = this.$refs.localAudio
      if (audio.webkitRequestFullscreen) {
        audio.webkitRequestFullscreen()
      } else if (audio.mozRequestFullScreen) {
        audio.mozRequestFullScreen()
      } else if (audio.requestFullscreen) {
        audio.requestFullscreen()
      }
    },

    getUserMedia(constraints, success, error) {
      if (navigator.mediaDevices.getUserMedia) {
        navigator.mediaDevices.getUserMedia(constraints).then(success).catch(error)
      } else if (navigator.webkitGetUserMedia) {
        navigator.webkitGetUserMedia(constraints).then(success).catch(error)
      } else if (navigator.mozGetUserMedia) {
        navigator.mozGetUserMedia(constraints).then(success).catch(error)
      } else if (navigator.getUserMedia) {
        navigator.getUserMedia(constraints).then(success).catch(error)
      }
    },
    startCall() {
      const remoteId = this.youPeerid
      if (remoteId === '') {
        alert('请输入对方ID')
        return
      }
      this.call = this.peer.call(remoteId, this.localStream)
      this.call.on('stream', (remoteStream) => {
        console.log(remoteStream)
        this.remoteStream = remoteStream
        // this.refs.remoteAudio.srcObject = remoteStream
        console.log(this.$refs.remoteAudio)
        console.log(this.$refs.remoteAudio.src)
        console.log(remoteStream)
        this.$refs.remoteAudio.srcObject = remoteStream
        this.$refs.remoteAudio.autoplay = true
        console.log(this.$refs.remoteAudio.src)
      })
      this.call.on('close', () => {
        console.log('用户通话关闭')
      })
      this.call.on('error', (err) => {
        console.log(err)
      })
    },
    endCall() {
      if (this.call) {
        this.call.close()
        this.call = null
      }
      this.remoteStream = null
      this.$refs.remoteAudio.src = null
      this.myPeerid = null

      this.localAudio = null
      this.refs.localAudio.srcObject = null
      this.arr = true
    }
  }
}
</script>
