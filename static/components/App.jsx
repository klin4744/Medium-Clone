import React from 'react'
import Navbar from './Navbar/Navbar'
import HomeHeader from './HomeHeader/HomeHeader.jsx'

class App extends React.Component {
    render() {
        return (
            <div>
                <Navbar />
                <HomeHeader />
            </div>
        )
    }
}

export default App
