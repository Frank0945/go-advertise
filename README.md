# go-advertise

An example of API service for creating and listing advertisements in GO.

## Project Layout

#### [Vertical slice architecture](https://www.jimmybogard.com/vertical-slice-architecture/)

此專案透過 Vertical slice 將相同功能的程式碼集中在同一個資料夾中，而不是按照它們的技術層級去分類。

## Technologies

### Dependency Injection

此專案使用 Uber 開發的依賴注入框架 [Fx](https://github.com/uber-go/fx) ，藉此方便的去抽換實作及管理依賴關係。

### Database Migration

在開發和維護中，資料庫的結構變化是很常見的需求。此專案透過 [go-migrate](https://github.com/golang-migrate/migrate) 來管理 schema 的版本。

### API Design

此專案透過 [Goa](https://github.com/goadesign/goa) 來進行 design-first 的 API 設計，開發時使用 `design.go` 來定義 API 的結構和 endpoints，然後生成相應的程式碼，包括路由、控制器、驗證和 docs 等。這種 design-first 的方法可以更好地設計和理解 API，並且減少重複的工作，提高開發效率。

## Unit Test

此專案目前採用了 [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) 來進行單元測試，然而在目前情境下，使用 fake database 可能更適合，因為它能驗證 SQL 的正確性。
