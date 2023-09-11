// SPDX-License-Identifier: Unlicense OR MIT

package gpu

import "gioui.org/gpu/internal/driver"

// An API carries the necessary GPU API specific resources to create a Device.
// There is an API type for each supported GPU API such as OpenGL and Direct3D.
type API = driver.API

// A RenderTarget denotes the destination framebuffer for a frame.
type RenderTarget = driver.RenderTarget

// OpenGLRenderTarget is a render target suitable for the OpenGL backend.
type OpenGLRenderTarget = driver.OpenGLRenderTarget

// Direct3D11RenderTarget is a render target suitable for the Direct3D 11 backend.
type Direct3D11RenderTarget = driver.Direct3D11RenderTarget

// MetalRenderTarget is a render target suitable for the Metal backend.
type MetalRenderTarget = driver.MetalRenderTarget

// VulkanRenderTarget is a render target suitable for the Vulkan backend.
type VulkanRenderTarget = driver.VulkanRenderTarget

// OpenGL denotes the OpenGL or OpenGL ES API.
type OpenGL = driver.OpenGL

// Direct3D11 denotes the Direct3D API.
type Direct3D11 = driver.Direct3D11

// Metal denotes the Apple Metal API.
type Metal = driver.Metal

// Vulkan denotes the Vulkan API.
type Vulkan = driver.Vulkan

// things for RenderOp
type Device = driver.Device
type FragmentShader = driver.FragmentShader
type VertexShader = driver.VertexShader
type Pipeline = driver.Pipeline
type PipelineDesc = driver.PipelineDesc
type VertexLayout = driver.VertexLayout
type InputDesc = driver.InputDesc
type Buffer = driver.Buffer
type Texture = driver.Texture
type LoadDesc = driver.LoadDesc
type BlendDesc = driver.BlendDesc

// ErrDeviceLost is returned from GPU operations when the underlying GPU device
// is lost and should be recreated.
var ErrDeviceLost = driver.ErrDeviceLost
