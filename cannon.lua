function Cannon(parent)
  local c = {}
  c.parent = parent
  c.x = parent.x
  c.y = parent.y
  c.angle = 0
  c.w = 20
  c.h = 30

  function c:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.rotate(self.angle-self.parent.angle)
      print(self.angle)
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
    lg.pop()
  end

  function c:update(dt)
    self.x = self.parent.x
    self.y = self.parent.y
    mouseX, mouseY = love.mouse.getPosition()
    self.angle = math.atan2((mouseY - self.y), (mouseX - self.x))
  end

  return c
end
