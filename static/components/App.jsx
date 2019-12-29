import React, {useState, useEffect} from 'react';
import Navbar from './Navbar/Navbar';
import HomeHeader from './HomeHeader/HomeHeader.jsx';
import HomeBottom from './HomeBottom/HomeBottom';
import axios from 'axios'

export default function App (){
    const [articles,setArticles] = useState([]);
    useEffect(() => {
        const getArticles = async () => {
            const { data: articles } = await axios.get('/api/articles');
            setArticles(articles);
        }
        if(!articles.length){
         getArticles();
        }
    }, [articles])
        console.log(articles)
      return (
         <div>
            <Navbar />
            <HomeHeader articles={articles} />
            <HomeBottom articles={articles} />
         </div>
      );

}
