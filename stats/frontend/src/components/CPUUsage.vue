<template>
  <div class="container">
    <h2>CPU Usage</h2>

 <vc-donut :auto-adjust-text-size="true" background="#0d1942" :thickness="10" :total="100" foreground="#515563" :sections="sections">
   <h1 style="color:#108ae3;">{{perc}}%</h1>
 </vc-donut>

  </div>
</template>

<script>
import * as Wails from '@wailsapp/runtime';

export default {

  data() {
    return {
      sections:[0],
      perc:""
    
    };
  },
  mounted:function(){
    Wails.Events.On("usage",usage=>{
      
      this.sections=[{value:usage.avg,color:'#108ae3'}]
      this.perc=usage.avg
    })

  },
  methods: {
  
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
