package nlp

import "sync"

// type CtxPool struct {
// 	p sync.Pool
// 	s sync.Pool
// }
//
// var (
// 	cp CtxPool
// )
//
// func (p *CtxPool) Get() *Ctx {
// 	v := p.p.Get()
// 	if v != nil {
// 		if c, ok := v.(*Ctx[T]); ok {
// 			return c
// 		}
// 	}
// 	return NewCtx[T]()
// }
//
// func (p *CtxPool[T]) Put(ctx *Ctx[T]) {
// 	ctx.Reset()
// 	p.p.Put(ctx)
// }

var cp sync.Pool

func AcquireCtx[T Byteseq]() *Ctx[T] {
	v := cp.Get()
	if v != nil {
		if ctx, ok := v.(*Ctx[T]); ok {
			return ctx
		}
	}
	return NewCtx[T]()
}

func ReleaseCtx[T Byteseq](ctx *Ctx[T]) {
	ctx.Reset()
	cp.Put(ctx)
}
