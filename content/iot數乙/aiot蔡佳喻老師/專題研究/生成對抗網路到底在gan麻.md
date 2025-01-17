[[專題研究]]
[[參考資料]]
https://medium.com/@hiskio/%E7%94%9F%E6%88%90%E5%B0%8D%E6%8A%97%E7%B6%B2%E8%B7%AF%E5%88%B0%E5%BA%95%E5%9C%A8gan%E9%BA%BB-f149efb9eb6b

要了解什麼是生成對抗網路，首先必須知道的是深度學習（Deep Learning）與機器學習（Machine Learning）還有人工智慧（Artificial Intelligence）之間的關係。以下是HiSKIO為這三種專有名詞的做的分類圖。

![](https://miro.medium.com/v2/resize:fit:700/1*gbpALjvkwADyTln2FN-Khg.png)

各用一句簡單的話來解釋三者：人工智慧是最早出現的概念，也是最終期望的成果 — — 讓機器具有如同人類甚至更多的思辨能力；機器學習則是能夠達成人工智慧的方法，透過與人類相似的學習方法，訓練機器進行資料分類、處理與預測；最後是深度學習，代表實現機器學習的一種技術。

深度學習中又以人工神經網路相關的應用最為熱門，而生成對抗網路（Generative Adversarial Network，以下簡稱GAN ）便是其中之一。

**_GAN的特別之處_**

GAN由兩個網路構成，分別是鑑別網路（Discriminating Network）與生成網路（Generative Network），透過兩者相互對抗產生結果是其深度學習的運作原理。簡單做個比喻：GAN是一場鑑定師與仿畫家的比賽，仿畫家畫出假畫讓鑑定師評斷有多接近真品，根據評斷結果再繼續畫出比原本更好的作品，鑑定師也會透過不斷練習提升鑑定水準，最後比賽的結果就是一幅幾可亂真的機器畫。

深度學習中還有其他許多神經網路，例如適合==處理空間資料==的卷積神經網路（Convoulutional Neural Network）、擅長處理時間序列與語意結構判斷的循環神經網路（Ruccurent Neural Network）等，而生成對抗網路也可以與前兩者結合（或者其他更多不同的網路），讓應用範圍加廣泛。具備其他神經網路沒有的「雙胞胎競爭」特性，也使它成為深度學習的一顆閃亮新星。 

**_深度學習的重要突破_**

以往訓練神經網路常是透過人類提供大量標記資料供機器分析練習（即監督式學習）比如知名的圍棋人工智慧AlphaGo前期的訓練，便是針對人類輸入的大量棋譜進行監督式學習（後期開始自我對弈的訓練則是非監督式學習，某種層面上也有點對抗訓練的概念）

而GAN透過自己相互對抗的生成與鑑別網路，大幅減少資料量的需求，也為非監督式學習提供了更為進步的方法。

深度學習領域的巨擘，同時也是Facebook的AI研究院長楊立昆（Yann LeCun）曾表示對抗訓練是有史以來最酷的技術（Adversarial training is the coolest thing since sliced bread），對於GAN他也說：「它為創建無監督學習模型提供了強有力的算法框架，有望幫助我們為AI加入常識（common sense）。我們認為，沿著這條路走下去，有不小的成功機會能開發出更智慧的AI。」

**_GAN的應用與展望_**

目前GAN較多被應用在生成資料方面，如圖像與影音的生成、合成、辨識、修復等等，進階一點的則是輸入文本描述便能生成與形容相符的圖像，或者透過語言模型實現機器翻譯等。由於GAN的熱門，一位叫Avinash Hindupur的工程師在GitHub建立了一份清單叫GAN動物園（the GAN zoo），整理了來自許多人研究各種GAN的論文，點[連結](https://github.com/hindupuravinash/the-gan-zoo)進去可以看到更多不同名字的GAN都被用來做什麼。

其他實際案例上，Apple公司透過改良的GAN，成功訓練出一個叫Refiner的網路，能生成更擬真的合成圖，降低標示訓練圖像識別樣本的成本；Nvidia公司的GAN則是將白天的街景轉為夜晚，作為自動駕駛車輛的訓練樣本，解決夜晚圖像資料不足的問題。

除了生成資料降低企業成本，作為非監督學習的重要訓練方法，儘管還有許多問題需要改良修正，GAN在未來依舊可能真正實現「完全不靠人類就能自主學習的AI」，其技術上的突破將會是人工智慧的發展趨勢。