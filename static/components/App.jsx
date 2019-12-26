import React from 'react';
import Navbar from './Navbar/Navbar';
import HomeHeader from './HomeHeader/HomeHeader.jsx';
import HomeBottom from './HomeBottom/HomeBottom';

class App extends React.Component {
   render() {
      return (
         <div>
            <Navbar />
            <HomeHeader />
            <HomeBottom />
         </div>
      );
   }
}

export default App;
