# cryptocurrency-price-forcaster
- This application was curated as Capstone project for obtaining my undergraduate degree in Computer Science. 
- The premise of the application was to utilize "Machine Learning" in a prescriptive fashion to ascertain future price predictions of Cryptocurrency assets. 
- The application is housed within this Mono Repo containing both the frontend and backend code.

## Core Compoments:
- The backend code is housing the primary algorithms utiilized for curating the price predictions. 
- The price predictions are curated using the mean, standard deviation, and normal distrubution of the dataset, which is the OHLC data of a selected crypocurrency. 
- This is performed via an implementation of the Naieve Gaussian Bayes algorithm. 
- The mathmatical equations used can be seen within the "algorithm" package within the Backend code.
- The mathmetical equatinos:
- ![Equation](./resources/images/maths.png)
