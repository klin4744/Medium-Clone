import React, { useState, useEffect } from 'react';
import Navbar from './Navbar/Navbar';
import HomeHeader from './HomeHeader/HomeHeader.jsx';
import HomeBottom from './HomeBottom/HomeBottom';
import SingleArticle from './SingleArticle/SingleArticle';
import { Route, BrowserRouter as Router, Switch, Link } from 'react-router-dom';
import axios from 'axios';

export default function App() {
   const [articles, setArticles] = useState([]);
   useEffect(() => {
      const getArticles = async () => {
         const { data: articles } = await axios.get('/api/articles');
         setArticles(articles);
      };
      if (!articles.length) {
         getArticles();
      }
   }, [articles]);
   return (
      <Router>
         <Navbar />
         <Switch>
            <Route
               exact
               path='/'
               render={navProps => (
                  <>
                     <HomeHeader {...navProps} articles={articles} />
                     <HomeBottom {...navProps} articles={articles} />
                  </>
               )}
            />
            <Route exact path='/articles/:id' component={SingleArticle} />
         </Switch>
         <Link to='/articles/1'>Go to 1</Link>
      </Router>
   );
}
