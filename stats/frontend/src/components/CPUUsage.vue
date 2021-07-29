<template>
  <div class="container">

 <vc-donut background="blue" :thickness="40" :total="100" foreground="grey" :sections="sections">
   <h1 style="color:rgb(16, 138, 227)">{{perc}}%</h1>
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
      console.log(usage.avg)
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
