
[[Kaggle]]
[[Python]]
[[aiot徐昌旭]]
[[aiot蔡佳喻]]

```python
from sklearn import preprocessing
from sklearn.model_selection import GridSearchCV
from sklearn.ensemble import RandomForestClassifier
from sklearn.ensemble import RandomForestRegressor

import warnings
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

%matplotlib inline

train = pd.read_csv("Titanic/train.csv")
test = pd.read_csv("Titanic/test.csv")
submit = pd.read_csv("Titanic/gender_submission.csv")

tranin.info()
test.info()
train.describe()
test.describe()
data = pd.concat([train, test])

data.reset_index(inplace=True,drop=True)

sns.countplot(x=data['Survived'])

sns.countplot(x=data['Pclass'],hue=data['Survived'])

sns.countplot(x=data['Sex'],hue=data['Survived'])
sns.countplot(x=data['Embarked'],hue=data['Survived'])

g=sns.FacetGrid(data,col='Survived')
g.map(sns.histplot,'Age',kde=False)

g=sns.FacetGrid(data,col='Survived')
g.map(sns.histplot,'Fare',kde=False)

g=sns.FacetGrid(data,col='Survived')
g.map(sns.histplot,'Parch',kde=False)

g=sns.FacetGrid(data,col='Survived')
g.map(sns.histplot,'SibSp',kde=False)

data['Family_Size'] = data['Parch']+data['SibSp']

g=sns.FacetGrid(data,col='Survived')
g.map(sns.histplot,'Family_Size',kde=False)

data['Title1'] = data['Name'].str.split(", ", expand=True)[1]

data['Name'].str.split(", ", expand=True).head(3)

data['Title1'].head(3)

data['Title1'] = data['Title1'].str.split(".", expand=True)[0]

data['Title1'].unique()

pd.crosstab(data['Title1'],data['Sex']).T.style.background_gradient(cmap='summer_r')

pd.crosstab(data['Title1'],data['Survived']).T.style.background_gradient(cmap='summer_r')


```

http://localhost:8889/notebooks/Desktop/01.%20Python/51.koggle/testkoggle.ipynb#

https://medium.com/jameslearningnote/%E8%B3%87%E6%96%99%E5%88%86%E6%9E%90-%E6%A9%9F%E5%99%A8%E5%AD%B8%E7%BF%92-%E7%AC%AC4-1%E8%AC%9B-kaggle%E7%AB%B6%E8%B3%BD-%E9%90%B5%E9%81%94%E5%B0%BC%E8%99%9F%E7%94%9F%E5%AD%98%E9%A0%90%E6%B8%AC-%E5%89%8D16-%E6%8E%92%E5%90%8D-a8842fea7077

載入隨機森林演算法(Random Forest)來預測存活率
去chatgpt

https://www.imagetotext.io/

https://zh.snipaste.com/

檢查一下年紀補資料的部分

![[Pasted image 20230926090827.png]]

太棒拉

