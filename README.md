# go-signalx

## Example

```go
go func() {
    select {
    case <-ctx.Done():
        zaplogger.Warn("context timed out", zap.Error(ctx.Err()))
    case s := <-signalx.QuitC():
        zaplogger.Sugar().Warnf("received quit signal %s", s)
    }
}()
```
