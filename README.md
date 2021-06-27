# Go API client for xrely

The go library for the xrely platform. For more information, please visit *https://www.xrely.com*

## Overview

- API version: 1.0.0
- Package version: 1.0.0

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./xrely"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.xrely.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CollectApi* | [**CollectActivityGet**](docs/CollectApi.md#collectactivityget) | **Get** /collect/activity | collects information related to user activity
*CollectApi* | [**CollectInfoGet**](docs/CollectApi.md#collectinfoget) | **Get** /collect/info | collects information related to device
*IndexApi* | [**IndexPost**](docs/IndexApi.md#indexpost) | **Post** /index | Insert data to your index bucket
*SearchApi* | [**SearchGet**](docs/SearchApi.md#searchget) | **Get** /search | Get autocomplete phrase or keywords for your query
*SearchApi* | [**SearchPost**](docs/SearchApi.md#searchpost) | **Post** /search | Provides relevant result for your query


## Documentation For Models

 - [AggrigationField](docs/AggrigationField.md)
 - [FilterField](docs/FilterField.md)
 - [IndexApiResponse](docs/IndexApiResponse.md)
 - [IndexMessage](docs/IndexMessage.md)
 - [ModelApiResponse](docs/ModelApiResponse.md)
 - [SearchRequest](docs/SearchRequest.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author
team xrely
info@xrely.com

