![Build status](https://github.com/evolidev/bus/actions/workflows/main.yml/badge.svg)
[![codecov](https://codecov.io/github/evolidev/bus/branch/main/graph/badge.svg?token=F1T68P1LVV)](https://codecov.io/github/evolidev/bus)

Bus Package
===========

Bus Struct
----------

The `Bus` struct is used to manage event-driven communications between objects in your application. The struct has the following fields:

*   `subscriber`: A slice of `Subscriber` objects, used to subscribe to events.
*   `pipelines`: A slice of `Pipeline` objects, used to filter events before they reach subscribers.
*   `dispatcher`: A `Dispatcher` object, used to manage the dispatching of events to subscribers.

NewBus Method
-------------

The `NewBus` method is used to create a new instance of the `Bus` struct. The method takes a variable number of `Pipeline` objects as parameters and returns a pointer to the newly created `Bus` object.

```go
func NewBus(pipelines ...Pipeline) *Bus
```


UseDispatcher Method
--------------------

The `UseDispatcher` method is used to set a custom `Dispatcher` for the `Bus`. The method returns a pointer to the `Bus` instance, allowing you to chain method calls.

```go
func (b *Bus) UseDispatcher(dispatcher Dispatcher) *Bus
```


Dispatch Method
---------------

The `Dispatch` method is used to dispatch an event to the `Bus`. The method takes an `EmitObject` as a parameter and dispatches it to all subscribers and pipelines.

```go
func (b *Bus) Dispatch(o EmitObject)
```

Subscribe Method
----------------

The `Subscribe` method is used to subscribe to events on the `Bus`. The method takes a `Subscriber` object as a parameter and adds it to the `subscriber` slice of the `Bus`.

```go
func (b *Bus) Subscribe(subscriber Subscriber)
```


PipeThrough Method
------------------

The `PipeThrough` method is used to add a new `Pipeline` to the `Bus`. The method takes a `Pipeline` object as a parameter and adds it to the `pipelines` slice of the `Bus`. The method returns a pointer to the `Bus` instance, allowing you to chain method calls.

```go
func (b *Bus) PipeThrough(pipe Pipeline) *Bus
```


PipeOnlyThrough Method
----------------------

The `PipeOnlyThrough` method is used to set the `Pipeline` for the `Bus`. The method takes a `Pipeline` object as a parameter and sets it as the only pipeline in the `pipelines` slice of the `Bus`. The method returns a pointer to the `Bus` instance, allowing you to chain method calls.

```go
func (b *Bus) PipeOnlyThrough(pipe Pipeline) *Bus
```
